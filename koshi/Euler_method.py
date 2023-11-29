import matplotlib.pyplot as plt
import numpy as np
import math

def f(x, y):
    return y * math.cos(x)


def method(f,n, x, y):
    for i in range(n):
        xn = x[i]+h
        yn = y[i] + h * f(x[i], y[i])
        x.append(xn)
        y.append(yn)
    return x, y
        
x=[0]
y=[1]
a, b = 0, 1
h = float(input("Введите h: "))
n= int((b-a)/h)
x, y = method(f, n,x,y)
for i in range (len(x)):
    print(i, ": " f"x[{i}] = {x[i]}, y[{i}] = {y[i]}")

x1 = np.arange(0, 1+h, h)
y1= np.exp(np.sin(x1))

plt.title("Метод Эйлера")
plt.xlabel("ось абсцисс")
plt.ylabel("ось ординат")
plt.plot(x, y)
plt.plot(x1,y1)
plt.legend(["Метод Эйлера","Точное решение"], loc = "center left")
plt.grid()
plt.show()