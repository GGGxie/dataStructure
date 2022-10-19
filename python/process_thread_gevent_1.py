"""10000个请求，开启2个进程，进程中开启3个线程，线程中开启5个协程来处理
"""

import requests, time
requests.packages.urllib3.disable_warnings()
from multiprocessing import Queue, Process
import threading
import gevent
import os
from gevent import monkey

# con1_2022-9-14_发现不能开猴子补丁,开了以后,运行1分多钟了,协程还没建立起来.
# monkey.patch_all()


def process_work(q, p_name):
    """
    创建3个线程
    :param q: 队列
    :param pname: 进程名
    :return:
    """
    thread_list = []
    for i in range(3):
        t_name = "{}--t--{}".format(p_name, i)
        t = threading.Thread(target=thread_work, args=(q, t_name))
        # print('创建线程---{}'.format(t_name))
        t.start()
        thread_list.append(t)
    for thread in thread_list:
        thread.join()


def thread_work(q, t_name):
    """
    创建5个协程
    :param q: 队列名
    :param tname: 线程名
    :return:
    """
    g_list = []
    for i in range(5):
        g_name = "{}--g--{}".format(t_name, i)  # 协程名
        # print("创建协程----{}".format(g_name))
        g = gevent.spawn(gevent_work, q, g_name)
        g_list.append(g)
    gevent.joinall(g_list)


def gevent_work(q, g_name):
    """
    协程做的事：处理任务
    :param q: 队列
    :param gname: 协程名
    :return:
    """
    count = 0
    while not q.empty():
        url = q.get(timeout=0.01)
        response = requests.get(url=url,verify=False)
        # print(response)
        # print(response.status_code)
        # con1_2022-9-14_打开下面的sleep开关,才有5个协程的切换,不然不切换,不然仅仅第一个协程处理任务,其余4个处理任务数为0.
        # 原来这里是gevent.sleep(0.01),把参数改为0.001,0.0001都可以切换.
        gevent.sleep(0.0001)
        count += 1
    print("----协程{}执行了{}个任务".format(g_name, count))


def count_time(old_func):
    """函数计时装饰器"""

    def wrapper(*args, **kwargs):
        print('开始执行')
        st = time.time()
        old_func(*args, **kwargs)
        et = time.time()
        print('结束执行')
        print('执行耗时：{}s'.format(et - st))

    return wrapper


@count_time
def main():
    """
    main函数创建2个进程,控制程序的运行
    :return:
    """
    q = Queue()
    # 创建10000个任务在队中
    # for i in range(10000):
    #     q.put("http://127.0.0.1:5000/demo")
    # print("11111")

    for i in range(400):
        # url = "https://python.freelycode.com/contribution/detail/{}".format(i)
        url = "https://news.cnblogs.com/n/page/{}".format(i)
        q.put(url)

    # print("q.qsize():",q.qsize())  #TODO执行就报错?
    process_list = []
    # 将创建的进程都加入进程列表，并启动
    n = os.cpu_count()
    for i in range(n):
        p_name = 'p--{}'.format(i)
        # print('创建进程-{}'.format(p_name))
        pro = Process(target=process_work, args=(q, p_name))  # 进程不共享全局变量，所有q做参数传进去
        process_list.append(pro)
        pro.start()
    # 主进程等待所有子进程
    for pro in process_list:
        pro.join()
if __name__ == '__main__':
    main()