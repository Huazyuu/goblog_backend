package randomname

type RandomType int

const (
	adjectiveAndPerson RandomType = iota // 形容词+人物 例如：帅气的罗纳尔多
	personActSomething                   // 人物+做事情 例如：梅西吃爆米花
)

// ActSomethingSlice 做什么事情
var actSomethingSlice = []string{
	"打豆豆", "吃爆米花", "走向人生巅峰", "大力出奇迹", "得到了金球奖", "爆射世界杯", "有亿点点忧伤", "完成了梅开二度", "完成了帽子戏法", "完成了大四喜", "完成了五子登科",
	"望穿秋水", "掐指一算", "横扫千军", "一眼定情", "会发财", "求而不得", "舞力四射", "横扫六合", "救了一个美女", "使出了佛怒火莲", "在巴黎徒步", "心花怒放", "在武汉看电影",
}
var actSomethingSliceCount = len(actSomethingSlice)

// adjectiveSlice 形容词
var adjectiveSlice = []string{
	"帅气的", "迷人的", "大汗淋漓的", "朝气蓬勃的", "隐身的", "天神下凡的", "快乐的", "冷静的", "醉熏的", "潇洒的", "糊涂的", "积极的",
	"冷酷的", "深情的", "粗暴的", "温柔的", "可爱的", "愉快的", "义气的", "认真的", "威武的", "传统的", "潇洒的", "漂亮的", "自然的",
	"专一的", "听话的", "昏睡的", "狂野的", "等待的", "搞怪的", "幽默的", "魁梧的", "活泼的", "开心的", "高兴的", "超帅的", "留胡子的",
	"坦率的", "直率的", "轻松的", "痴情的", "完美的", "精明的", "无聊的", "有魅力的", "丰富的", "高挑的", "傻傻的", "冷艳的", "爱听歌的",
	"繁荣的", "饱满的", "炙热的", "暴躁的", "碧蓝的", "俊逸的", "英勇的", "健忘的", "故意的", "无心的", "土豪的", "朴实的", "兴奋的",
	"幸福的", "淡定的", "不安的", "阔达的", "孤独的", "独特的", "疯狂的", "时尚的", "落后的", "风趣的", "忧伤的", "大胆的", "爱笑的", "矮小的",
	"健康的", "合适的", "玩命的", "沉默的", "斯文的", "香蕉的", "苹果的", "鲤鱼的", "鳗鱼的", "任性的", "细心的", "粗心的", "大意的", "甜甜的",
	"酷酷的", "健壮的", "英俊的", "霸气的", "阳光的", "默默的", "大力的", "孝顺的", "忧虑的", "着急的", "紧张的", "善良的", "凶狠的", "害怕的",
	"重要的", "危机的", "欢喜的", "欣慰的", "满意的", "跳跃的", "诚心的", "称心的", "如意的", "怡然的", "娇气的", "无奈的", "无语的", "激动的",
	"愤怒的", "美好的", "感动的", "激情的", "激昂的", "震动的", "虚拟的", "超级的", "寒冷的", "精明的", "明理的", "犹豫的", "忧郁的", "寂寞的",
	"奋斗的", "勤奋的", "现代的", "过时的", "稳重的", "热情的", "含蓄的", "开放的", "无辜的", "多情的", "纯真的", "拉长的", "热心的", "从容的",
	"体贴的", "风中的", "曾经的", "追寻的", "儒雅的", "优雅的", "开朗的", "外向的", "内向的", "清爽的", "文艺的", "长情的", "平常的", "单身的",
	"伶俐的", "高大的", "懦弱的", "柔弱的", "爱笑的", "乐观的", "耍酷的", "酷炫的", "神勇的", "年轻的", "唠叨的", "瘦瘦的", "无情的", "包容的",
	"顺心的", "畅快的", "舒适的", "靓丽的", "负责的", "背后的", "简单的", "谦让的", "彩色的", "缥缈的", "欢呼的", "生动的", "复杂的", "慈祥的",
	"仁爱的", "魔幻的", "虚幻的", "淡然的", "受伤的", "雪白的", "高高的", "糟糕的", "顺利的", "闪闪的", "羞涩的", "缓慢的", "迅速的", "优秀的",
	"聪明的", "含糊的", "俏皮的", "淡淡的", "坚强的", "平淡的", "欣喜的", "能干的", "灵巧的", "友好的", "机智的", "机灵的", "正直的", "谨慎的",
	"俭朴的", "殷勤的", "虚心的", "辛勤的", "自觉的", "无私的", "无限的", "踏实的", "老实的", "现实的", "可靠的", "务实的", "拼搏的", "个性的",
	"粗犷的", "活力的", "成就的", "勤劳的", "单纯的", "落寞的", "朴素的", "悲凉的", "忧心的", "洁净的", "清秀的", "自由的", "小巧的", "单薄的",
	"贪玩的", "刻苦的", "干净的", "壮观的", "和谐的", "文静的", "调皮的", "害羞的", "安详的", "自信的", "端庄的", "坚定的", "美满的", "舒心的",
	"温暖的", "专注的", "勤恳的", "美丽的", "腼腆的", "优美的", "甜美的", "甜蜜的", "整齐的", "动人的", "典雅的", "尊敬的", "舒服的", "妩媚的",
	"秀丽的", "喜悦的", "甜美的", "彪壮的", "强健的", "大方的", "俊秀的", "聪慧的", "陶醉的", "悦耳的", "动听的", "明亮的", "结实的", "魁梧的",
	"标致的", "清脆的", "敏感的", "光亮的", "大气的", "老迟到的", "知性的", "冷傲的", "呆萌的", "野性的", "隐形的", "笑点低的", "微笑的", "笨笨的",
	"难过的", "沉静的", "火星上的", "失眠的", "安静的", "纯情的", "要减肥的", "迷路的", "烂漫的", "哭泣的", "贤惠的", "苗条的", "温婉的", "发嗲的",
	"会撒娇的", "贪玩的", "执着的", "眯眯眼的", "花痴的", "想人陪的", "眼睛大的", "高贵的", "傲娇的", "心灵美的", "爱撒娇的", "细腻的", "天真的",
	"怕黑的", "感性的", "飘逸的", "怕孤独的", "忐忑的", "还单身的", "怕孤单的", "懵懂的",
}
var adjectiveSliceCount = len(adjectiveSlice)

// personSlice 人物
var personSlice = []string{
	"梅西", "罗纳尔多", "内马尔", "贝克汉姆", "姆巴佩", "苏牙", "小罗", "大罗", "鲁尼", "本泽马", "贝利", "约翰", "柯南", "米老鼠",
	"弗朗茨", "普拉蒂尼", "迪斯蒂法诺", "普斯卡什", "乔治·贝斯特", "范巴斯滕", "尤西比奥", "列夫·雅辛", "博比", "博比·摩尔", "盖德·穆勒",
	"巴乔", "斯坦利", "济科", "弗兰科", "加林查", "保罗", "肯尼", "巴蒂斯图塔", "坎通纳", "哈吉", "罗马里奥",
	"齐达內", "古利特", "约翰·查尔斯", "卡卡", "哈维", "伊布", "莫德里奇", "欧文", "巴乔", "萧炎", "鸣人", "佐助", "盖茨比", "乔布斯", "索尔",
	"宙斯", "雅典娜", "星矢", "托尼", "科比", "乔丹", "西瓜太郎", "一休", "莫甘娜", "永恩", "爱因斯坦", "肖申克", "摩尔",
}
var PersonSliceCount = len(personSlice)
