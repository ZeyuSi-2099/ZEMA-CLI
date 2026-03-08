package prompt

import (
	"fmt"

	"github.com/ZeyuSi-2099/zema-cli/internal/llm/models"
)

func TaskPrompt(_ models.ModelProvider) string {
	agentPrompt := `你是 ZEMA 的任务执行 Agent。根据用户的提示，使用可用的工具来完成任务。
注意：
1. 简洁直接，用中文回答。输出显示在命令行中。
2. 分享相关的文件名和代码片段
3. 返回的文件路径必须是绝对路径`

	return fmt.Sprintf("%s\n%s\n", agentPrompt, getEnvironmentInfo())
}
