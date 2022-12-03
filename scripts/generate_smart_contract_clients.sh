cd "$(dirname "$0")" # Set current working directory to always be /scripts
abigen --pkg position_manager_contract --abi ../internal/adapters/smart_contracts/abi/position_manager_contract/position_manager.json --out ../internal/adapters/smart_contracts/abi/position_manager_contract/position_manager.go
abigen --pkg vault_contract --abi ../internal/adapters/smart_contracts/abi/vault_contract/vault.json --out ../internal/adapters/smart_contracts/abi/vault_contract/vault.go
