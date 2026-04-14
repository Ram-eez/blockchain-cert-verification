const hre = require("hardhat");

async function main() {
  const Registry = await hre.ethers.getContractFactory("CertificateRegistry");
  const registry = await Registry.deploy();
  await registry.waitForDeployment();

  const address = await registry.getAddress();
  console.log("CertificateRegistry deployed to:", address);

  const fs = require("fs");
  fs.writeFileSync(
    "deployments.json",
    JSON.stringify({ CertificateRegistry: address, network: "sepolia" }, null, 2)
  );
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});