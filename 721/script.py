from collections import defaultdict
class Solution:
    def accountsMerge(self, accounts):
        names = {}
        graph = defaultdict(set)
        for acc in accounts:
            name = acc[0]
            for email in acc[1:]:
                graph[acc[1]].add(email)
                graph[email].add(acc[1])
                names[email] = name

        comps, seen, ans, i = defaultdict(list), set(), [], 0
        def dfs(node, i):
            comps[i].append(node)
            seen.add(node)
            for neib in graph[node]:
                if neib not in seen: dfs(neib, i)

        for email in graph:
            if email not in seen:
                dfs(email, i)
                i += 1

        return [[names[val[0]]] + sorted(val) for _,val in comps.items()]

if __name__=='__main__':
    obj = Solution()
    val = [["John","johnsmith@mail.com","john_newyork@mail.com"],["John","johnsmith@mail.com","john00@mail.com"],["Mary","mary@mail.com"],["John","johnnybravo@mail.com"]]
    result = obj.accountsMerge(val)
    print(result)
