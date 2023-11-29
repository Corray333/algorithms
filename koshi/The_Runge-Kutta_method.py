import matplotlib.pyplot as plt
import numpy as np
import math

def f(x, y):
    return y * math.cos(x)


def method(f,n, x, y, h):
    for i in range(n):
        k1 = f(x[i], y[i])
        k2 = f(x[i]+h/2, y[i]+(h*k1)/2)
        k3 = f(x[i]+h/2, y[i]+(h*k2)/2)
        k4 = f(x[i]+h, y[i]+h*k3)
        yn = y[i] + h/6 * (k1 + 2*k2 + 2*k3 + k4)
        xn = x[i]+h
        x.append(xn)
        y.append(yn)
    return x, y
        
x=[0]
y=[1]
a, b = 0, 1
h = float(input("Введите h: "))
n= int((b-a)/h)
x, y = method(f, n,x,y, h)
for i in range (len(x)):
    print(i, ": " f"x[{i}] = {x[i]}, y[{i}] = {y[i]}")

x1 = np.arange(0, 1+h, h)
y1= np.exp(np.sin(x1))


plt.title("Метод Рунге-Кутта 4-го порядка")
plt.xlabel("ось абсцисс")
plt.ylabel("ось ординат")
plt.plot(x, y)
plt.plot(x1, y1)
plt.legend(["Метод Рунге-Кутта 4-го порядка","Точное решение"], loc = "center left")
plt.grid()
plt.show()