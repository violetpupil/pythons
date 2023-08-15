# [pandas](https://pandas.pydata.org/)

```python
import matplotlib.pyplot as plt
import pandas as pd

# 读取csv文件
df = pd.read_csv("tmp.csv")
# 获取total列Series
column = df["total"]
# 分段点 最小值 三等分点 最大值
bins = [column.min(), column.quantile(0.33), column.quantile(0.66), column.max()]
# 分段 包含第一个分段点
segments = pd.cut(column, bins, include_lowest=True)
# 统计每个分段值的数量
counts = pd.value_counts(segments, sort=False)

# 柱状图 分段范围为x坐标 分段值数量为y坐标
b = plt.bar(counts.index.astype(str), counts)
# 添加数据标签
plt.bar_label(b, counts)
# 阻塞显示图表
plt.show()
```

```python
# 计算total列平均值
df["total"].mean()
# 生成total列统计信息
df["total"].describe()
```
