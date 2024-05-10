import numpy as np
vector1 = np.array([1,2,3])
vector2 = np.array([4,5,7])
manhaton_dist = np.sum(np.abs(vector1-vector2))
print("曼哈顿距离为", manhaton_dist)