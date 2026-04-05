const { ethers } = require("hardhat");

async function main() {
  const Certificate = await ethers.getContractFactory("Certificate");

  const contract = await Certificate.deploy();

  await contract.waitForDeployment();

  console.log("Deployed to:", contract.target);
}

main().catch((error) => {
  console.error(error);
  process.exit(1);
});