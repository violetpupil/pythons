# [matplotlib](https://matplotlib.org/)

```python
import matplotlib.pyplot as plt
```

```python
# 柱状图，x坐标，y坐标
num = [1, 2, 3, 4]
b = plt.bar(range(len(num)), num)
# 添加数据标签
plt.bar_label(b, num)
# 设置标题
plt.title("total")
# 阻塞显示图表
plt.show()
# 保存图片
plt.savefig("total.png")
```
