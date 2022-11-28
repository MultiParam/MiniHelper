import os

def whole_name(d):
    ld   = os.listdir(d)
    return list(map(lambda x : os.path.join(base, x), ld))

def filter_handler(x, include):
    if os.path.isfile(x) and len(os.path.basename(x).split('.')) > 1 and os.path.basename(x).split('.')[1] in include:
        return x

def walk_4_files(base, include):
    ld    = os.listdir(base)
    maps  = list(map(lambda x : os.path.join(base, x), ld))
    files = list(filter(lambda x : filter_handler(x, include), maps))
    dirs  = filter(lambda x : os.path.isdir(x), maps)
    for d in dirs:
        files += walk_4_files(d, include)
    return files