package structrue

//手写计算器III

func calculate(s string) int {
    n:=len(s)
    var dfs func(int)(int,int)
    dfs = func(i int)(int,int){
        num:=0
        preflag:='+'
        stack:=[]int{}
        for i<n{
            if '0' <= s[i] && s[i] <='9'{
                num = 0
                for i<n && '0' <= s[i] && s[i] <='9'{
                    num = num * 10 + int(s[i]-'0')
                    i++
                }
                i--
            }

            if s[i] == '('{
                num,i = dfs(i+1)
                i++
            }

            if i >=n-1 || (('0'>s[i] || s[i]>'9') && s[i]!=' '){
                switch preflag{
                case '+':
                    stack = append(stack , num)
                case '-':
                    stack = append(stack,-num)
                case '*':
                    stack[len(stack)-1]*=num
                case '/':
                    stack[len(stack)-1]/=num 
                }
                if i<n{
                    preflag = rune(s[i])
                }
            }

            if i >=n-1 || s[i]==')'{
                break
            }
            i++
        }
         ans:=0
            for _,v:=range stack{
                ans+=v
            }
            return ans ,i
    }
    sum,_:=dfs(0)
    return sum
}