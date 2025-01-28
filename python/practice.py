import numpy as np
import matplotlib.pyplot as plt
from IPython.display import display


def true_fun(X):  # 这是我们设定的真实函数，即ground truth的模型
    return 1.5 * X + 0.2


np.random.seed(0)  # 设置随机种子
n_samples = 30  # 设置采样数据点的个数

"""生成随机数据作为训练集，并且加一些噪声"""
X_train = np.sort(np.random.rand(n_samples))
y_train = (true_fun(X_train) + np.random.randn(n_samples) * 0.05).reshape(n_samples, 1)
display(X_train)
print("11111")
display(y_train)
