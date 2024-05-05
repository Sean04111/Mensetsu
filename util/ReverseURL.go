package util



//反转 url 空间复杂度 O(1)
func ReverseURL(url []byte)string{
	l,r:=0,len(url)-1
	for i:=0;i<len(url);i++{
		if i+1 < len(url) && url[i]=='/' && url [i+1]=='/'{
			l = i + 2
			break
		}
	}

	ReverseBetween(url,l,r)

	for i,j:=l,l;j<len(url);j++{
		if 'a'<=url[j] && url[j]<='z'{
			continue
		}else{
			ReverseBetween(url,i,j-1)
			i = j+1
		}
	}
	return string(url)
}
func ReverseBetween(s []byte,from ,to int){
	for i,j:=from,to;i<=j;i,j = i+1,j-1{
		s[i],s[j] = s[j],s[i]
	}
}