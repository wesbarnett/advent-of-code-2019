from argparse import ArgumentParser
from pathlib import Path
from collections import defaultdict, deque

def count_orbits_recursive(objects):

    def traverse(node, count=0):
        if not objects[node]:
            return count
        orbits = count
        for obj in objects[node]:
            orbits += traverse(obj, count+1)
        return orbits

    return traverse("COM")

def count_orbits(objects):
    objects_deque = deque(["COM"])
    count = 0
    tmp_count = 0
    while objects_deque:
        tmp_count += 1
        for obj in list(objects_deque):
            objects_deque.popleft()
            for x in objects[obj]:
                count += tmp_count
                objects_deque.append(x)

    return count


if __name__ == "__main__":
    parser = ArgumentParser()
    parser.add_argument("-i", "--infile", required=False, type=Path, default=Path("input"))
    args = parser.parse_args()

    lines = Path(args.infile).read_text().rstrip("\n").split("\n")

    objects = defaultdict(list)
    for line in lines:
        x, y = line.split(")")
        objects[x].append(y)

    print(count_orbits(objects))
