// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/* 
创建一个简单的NFT市场合约：

定义NFT结构体（id、owner、price、forSale）
铸造NFT功能
上架/下架功能
购买功能
查询所有在售NFT

提示：
使用ID自增模式
使用mapping存储NFT
使用array追踪在售列表
*/
contract NFTMarket {
    struct NFTToken {
        uint256 id;
        address owner;
        uint256 price;
        bool forSale;
    }
    mapping(uint256 => NFTToken) public nftTokens;
    NFTToken[] public NFTTokens;
    uint256 public nftTokenMaxID = 0;

    event NFTMint(uint256 id, address owner, uint256 price);
    event NFTOnSale(uint256 id);
    event NFTUnOnSale(uint256 id);
    event NFTBuy(uint256 id, address from, address to);

    // 铸造新的NFT
    function mintNFT(uint256 price) public returns (uint256) {
        // 是先把0赋值给tokenID，然后再 增加1
        uint tokenID = nftTokenMaxID++;
        NFTToken memory newNFT = NFTToken({
            id: tokenID,
            owner: msg.sender,
            price: price,
            forSale: true
        });
        nftTokens[tokenID] = newNFT;
        NFTTokens.push(newNFT);
        emit NFTMint(tokenID, msg.sender, price);
        return tokenID;
    }

    // 上架NFT
    function onSaleNFT(uint256 tokenID) public {
        require(tokenID < nftTokenMaxID, "Invalid token ID");
        require(!nftTokens[tokenID].forSale, "NFT already on sale");
        require(nftTokens[tokenID].owner == msg.sender, "Not owner");
        nftTokens[tokenID].forSale = true;
        NFTTokens[tokenID].forSale = true;
        emit NFTOnSale(tokenID);
    }

    // 下架NFT
    function unSaleNFT(uint256 tokenID) public {
        require(tokenID < nftTokenMaxID, "Invalid token ID");
        require(nftTokens[tokenID].forSale, "NFT not on sale");
        require(nftTokens[tokenID].owner == msg.sender, "Not owner");
        nftTokens[tokenID].forSale = false;
        NFTTokens[tokenID].forSale = false;
        emit NFTUnOnSale(tokenID);
    }
     
    // 这个函数还没调用过
    function buyNFT(uint256 tokenID) public payable {
        require(nftTokens[tokenID].forSale, "NFT not on sale");
        require(msg.value == nftTokens[tokenID].price, "Price not match");
        
        // 保存旧所有者地址
        address oldOwner = nftTokens[tokenID].owner;
        
        // 转账给卖家（使用 call 替代 transfer）
        (bool success, ) = payable(oldOwner).call{value: msg.value}("");
        require(success, "Transfer failed");
        require(address(this).balance == 0, "Contract should have zero balance after transfer");
        
        // 更新所有者和状态
        nftTokens[tokenID].owner = msg.sender;
        nftTokens[tokenID].forSale = false;
        
        emit NFTBuy(tokenID, oldOwner, msg.sender);
    }

    function getOnSaleNFTs() public view returns (NFTToken[] memory) {
        // 先统计在售NFT数量
        uint256 count = 0;
        for (uint256 i = 0; i < NFTTokens.length; i++) {
            if (NFTTokens[i].forSale) {
                count++;
            }
        }
        
        // 创建固定大小的memory数组
        NFTToken[] memory onSaleNFTs = new NFTToken[](count);
        
        // 填充数组
        uint256 index = 0;
        for (uint256 i = 0; i < NFTTokens.length; i++) {
            if (NFTTokens[i].forSale) {
                onSaleNFTs[index] = NFTTokens[i];
                index++;
            }
        }
        
        return onSaleNFTs;
    }
}