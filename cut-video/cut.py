import subprocess
import os
from files import walk_4_files


# out = subprocess.check_output(['ls']).decode('utf-8')
# print(out)

def get_video_duration(video):
    ffprobe_cmd = f'ffprobe -i {video} -show_entries format=duration -v quiet -of csv="p=0"'
    out = subprocess.check_output(ffprobe_cmd, shell=True).decode('utf-8').strip()
    print(f'{int(float(out))} s')
    return int(float(out))


def cut(video_path, output_path, include, delta_x):
    videos = walk_4_files(video_path, include) 
    delta_x_origin = delta_x
    for v in videos:
        name_without_ext = os.path.basename(v).split('.')[0]
        duration = get_video_duration(v)
        for i in range(0, duration + 1, delta_x_origin):
            delta_x   = delta_x_origin
            from_h    = i // 3600
            from_mins = i % 3600 // 60
            from_sec  = i % 60
            # 最后一段
            if i + delta_x > duration:
                delta_x = duration - i
            outfile = os.path.join(output_path, f'cut_{i}_from_{name_without_ext}.mp4')
            # to_mins = (i + delta_x) // 60
            # to_h    = from_mins // 60
            # to_sec  = (i + delta_x) % 60
            # cut_cmd = f'ffmpeg -i {v} -ss 00:{from_mins}:{from_sec} -to 00:{to_mins}:{to_sec} -strict -2 /mnt/shared/random/xdd/data/target_video/real/video/720/big/test/{i}.mp4'
            #! -ss 放在前头 -t 
            cut_cmd = f'ffmpeg -ss {from_h:02d}:{from_mins:02d}:{from_sec:02d} -t {delta_x} -y -i {v} -c copy {outfile}'
            subprocess.check_output(cut_cmd, shell=True)


# video_path = '/mnt/shared/random/xdd/data/target_video/real/video/720-big'
# output_path = '/mnt/shared/random/xdd/data/target_video/real/video/720-big/cut'
video_path  = './cut-source'
output_path = '/mnt/shared/random/xdd/data/target_video/real/video/1080-cut-10s'
delta_x    = 10 # s
#! ls -t | cut -d '_' -f 4 | uniq 检查一下切割了多少份视频
cut(video_path, output_path, ('mp4'), delta_x)

