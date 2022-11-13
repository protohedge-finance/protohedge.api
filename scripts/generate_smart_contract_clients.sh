cd "$(dirname "$0")" # Set current working directory to always be /scripts
abigen --pkg position_manager_contract --abi ../internal/contracts/position_manager_contract/position_manager.json --out ../internal/contracts/position_manager_contract/position_manager.go
abigen --pkg vault_contract --abi ../internal/contracts/vault_contract/vault.json --out ../internal/contracts/vault_contract/vault.go
