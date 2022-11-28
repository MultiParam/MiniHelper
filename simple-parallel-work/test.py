import time
import os
import subprocess
import logging
from watchdog.observers import Observer
from watchdog.events import LoggingEventHandler
import datetime
import time
from watchdog.observers import Observer
from watchdog.events import FileSystemEventHandler
from multiprocessing import Process,JoinableQueue,Manager

class ExtractFrame(Process): 
    def __init__(self, q):
        super().__init__()
        self.q = q
    def run(self):
        while True:
            if not self.q.empty():
                time.sleep(5)
                src = self.q.get()
                subprocess.call(f"cd ../extract_frame;bash ./xi_frames_single.sh {os.path.basename(src)}", shell=True)
                self.q.task_done()
 
class MyEventHandler(FileSystemEventHandler):
    def __init__(self, q):
        FileSystemEventHandler.__init__(self)
        self.q = q
 
    def on_any_event(self, event):
        pass
        # print("-----")
        # print(datetime.datetime.now().strftime('%Y-%m-%d %H:%M:%S.%f'))
 
    # 新建
    def on_created(self, event):
        if event.is_directory:
            # print("目录 created:{file_path}".format(file_path=event.src_path))
            pass
        else:
            print("created {file_path}".format(file_path=event.src_path))
 
    # 修改
    def on_modified(self, event):
        if event.is_directory:
            # print("modified:{file_path}".format(file_path=event.src_path))
            pass
        else:
            print("modified:{file_path}".format(file_path=event.src_path))
            # parent = os.path.realpath(os.path.join(event.src_path, ".."))
            self.q.put(event.src_path)
            q.join()
 
 
if __name__ == '__main__':
    logging.basicConfig(level=logging.INFO,
                        format='%(asctime)s - %(message)s',
                        datefmt='%Y-%m-%d %H:%M:%S')
 
    path = "/mnt/shared/random/xdd/data/target_video/fake/wav2lib/video/1080-hq-cut-10s"
    # path = "./"
 
    # 内置的LoggingEventHandler
    event_handler = LoggingEventHandler()
 
    # 观察者
    observer = Observer()

    # q = Manager().Queue() 
    q        = JoinableQueue()
    e        = ExtractFrame(q)
    e.daemon = True
    e.start()
    myEventHandler = MyEventHandler(q)
    # recursive:True 递归的检测文件夹下所有文件变化。
    observer.schedule(myEventHandler, path, recursive=True)
 
    # 观察线程，非阻塞式的。
    observer.start()
 
    try:
        while True:
            time.sleep(300)
    except KeyboardInterrupt:
        observer.stop()
    e.join()
    observer.join()