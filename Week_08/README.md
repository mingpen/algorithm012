# 146. LRU缓存机制

LRU :Least Recently Used.Discards the least recently used items first. This algorithm requires keeping track of what was used when, which is expensive if one wants to make sure the algorithm always discards the least recently used item.

- 最近操作的数据最年轻【写入、读取】；
- 缓存的数据容量有限；
- hashMap+linked List;
- hashMap 的指向 list中元素；
- 使用list顺序表示顺序【map没有顺序，没法保存新旧】;
- 写操作时，先通过map查看是否是否存在这个elemet,
  - 如果存在将其移动到front，修改值
  - 如果不存在：a. 没有超过容量，加到front【表示最新数据】,key加入到map ；b.超过容量，将list的rear的key从map中删除，并移为front，修改element 的key,value。
- list 的element中要存放key和value，因为移除老数据，需要从map中移除对应的key。

# 191. 位1的个数

        num &= num-1

将非0 最底位，变成0

# 52. N皇后 II

复习回溯算法
