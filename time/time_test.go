package time

import (
	"testing"
	"time"
)

func TestTimeNow(t *testing.T) {
	now := time.Now()
	t.Errorf("current time: [%v]\n", now)

	year := now.Year()
	month := now.Month()
	
	t.Errorf("year: %v, month: %v\n", year, month)
}

func TestTimezoneAndDate(t *testing.T) {
	// 中国没有夏令时，使用一个固定的8小时的UTC时差(东八区，UTC +08:00)，对于很多其他国家需要考虑夏令时。
	timezone := int((8 * time.Hour).Seconds())

	// FixedZone 返回始终使用给定区域名称和偏移量(UTC 以东秒)的Location。UTC +08:00
	shanghaiTimezone := time.FixedZone("Asia/Shanghai", timezone)

	// 如果当前系统有时区数据库，则可以加载一个位置得到对应的时区，例如，加载纽约所在的时区，UTC -05:00
	newYorkTimezone, _ := time.LoadLocation("America/New_York")

	utc := time.Date(2009, 1, 1, 12, 0, 0, 0, time.UTC)
	shanghai := time.Date(2009, 1, 1, 20, 0, 0, 0, shanghaiTimezone)
	NewYork := time.Date(2009, 1, 1, 7, 0, 0, 0, newYorkTimezone)

	// 北京时间（东八区）比UTC早8小时，所以上面两个时间看似差了8小时，但表示的是同一个时间
	t1 := utc.Equal(shanghai)

	// 纽约（西五区）比UTC晚5小时，所以上面两个时间看似差了5小时，但表示的是同一个时间
	t2  := utc.Equal(NewYork)


	t.Errorf("[%v] = [%v] => [%t]\n",utc,shanghai,t1)
	t.Errorf("[%v] = [%v] => [%t]\n",utc,NewYork,t2)
}

func TestTimestamp(t *testing.T) {
	now := time.Now()

	// 秒级时间戳
	timestamp := now.Unix()

	// 毫秒时间戳
	milli := now.UnixMilli()

	// 微秒级时间戳
	micro := now.UnixMicro()

	// 纳秒级时间戳
	nano := now.UnixNano()

	t.Errorf("\n秒级时间戳: %v\n毫级时间戳: %v\n微级时间戳: %v\n纳级时间戳: %v\n", timestamp, milli, micro, nano)
}

func TestTick(t *testing.T) {
	ticker := time.Tick(time.Second)

	for i := range ticker {
		t.Error(i)
	}
}

func TestTimeFormat(t *testing.T) {
	// 获取当前时间对象，后续方便基于对时间对象进行格式化操作
	now := time.Now()

	// 格式化的模板为 2006-01-02 15:04:05

	// 24小时制
	t.Error(now.Format("2006-01-02 15:04:05.0 Jan Mon"))

	// 12小时制
	t.Error(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))

	// 小数点后写0，因为有3个0所以格式化输出的结果也保留3位小数
	t.Error(now.Format("2006/01/02 15:04:05.000"))

	// 小数点后写9，会省略末尾可能出现的0,例如.910->.91
	t.Error(now.Format("2006/01/02 15:04:05.999"))

	// 只格式化时分秒部分
	t.Error(now.Format("15:04:05"))

	// 只格式化日期部分
	t.Error(now.Format("2006.1.2"))

	// 只格式化日期部分
	t.Error(now.Format("2006.01.02"))
}

func TestTimeParse(t *testing.T) {
	// 在没有时区指示符的情况下，time.Parse 返回UTC时间
	t1, _ := time.Parse("2006/01/02 15:04:05", "2030/10/05 11:25:20")

	// 在有时区指示符的情况下，time.Parse 返回对应时区的时间表示
	// RFC3339     = "2006-01-02T15:04:05Z07:00"
	t2, _ := time.Parse(time.RFC3339, "2030-10-05T11:25:20+08:00")

	t.Errorf("t1 = [%v], t2 = [%v]\n", t1, t2)
}

func TestParseInLocation(t *testing.T) {
	now := time.Now()

	// 加载时区
	timezone, _ := time.LoadLocation("Asia/Shanghai")

	// 按照指定时区和指定格式解析字符串时间
	t1, _ := time.ParseInLocation("2006/01/02 15:04:05", "2030/07/20 11:25:20", timezone)

	t.Error(now)

	t.Error(t1.Sub(now))
}

func TestTimeSince(t *testing.T) {
	start := time.Now()

	t.Error("执行代码")

	time.Sleep(2 * time.Second)

	t.Error("耗时", time.Since(start).Milliseconds(), "ms")
}