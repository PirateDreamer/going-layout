# 项目名
APP_NAME=going-layout

# Swagger 文件生成目录
SWAGGER_DIR=./doc/swagger
SWAGGER_FILE=$(SWAGGER_DIR)/swagger.json

# Apifox 配置
APIFOX_PROJECT_ID=6922858
APIFOX_API_KEY=APS-blmZxBSzo9TliE9n3bet0BErThlLzAmm
APIFOX_UPLOAD_URL=https://api.apifox.com/v1/projects//$(APIFOX_PROJECT_ID)/import-openapi?locale=zh-CN

# 生成 swagger 文档
swagger:
	@echo ">>> Generating Swagger documentation..."
	@swag init --parseDependency --parseInternal --output $(SWAGGER_DIR)
	@echo "Swagger documentation has been generated to $(SWAGGER_FILE)"

API_VERSION=2024-03-28

# 上传到 Apifox# 上传到 Apifox
apifox-upload:
	@echo ">>> Uploading Swagger documentation to Apifox..."
	@powershell -Command "& {$$content = Get-Content -Raw -Path $(SWAGGER_FILE); $$content = $$content -replace '\"', '\\\"'; $$body = @{input=$$content; options=@{targetEndpointFolderId=0; endpointOverwriteBehavior='OVERWRITE_EXISTING'; endpointCaseOverwriteBehavior='OVERWRITE_EXISTING'; updateFolderOfChangedEndpoint=$$false}} | ConvertTo-Json -Compress; $$response = Invoke-RestMethod -Uri '$(APIFOX_UPLOAD_URL)' -Method Post -Headers @{ 'X-Apifox-Api-Version' = '$(API_VERSION)'; 'Authorization' = 'Bearer $(APIFOX_API_KEY)'; 'Content-Type' = 'application/json' } -Body $$body; Write-Output $$response}"
	@echo ">>> Upload completed!"

# 一键生成并上传
gen-api: swagger apifox-upload
	@echo ">>> Swagger has been updated to Apifox ✅"
