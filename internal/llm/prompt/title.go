package prompt

import "github.com/ZeyuSi-2099/zema-cli/internal/llm/models"

func TitlePrompt(_ models.ModelProvider) string {
	return `根据用户的第一条消息生成一个简短的中文标题
- 不超过 25 个汉字
- 标题应概括用户消息的核心内容
- 只有一行
- 不要使用引号或冒号
- 你返回的全部文本将直接作为标题
- 永远不要返回超过一行的内容`
}
