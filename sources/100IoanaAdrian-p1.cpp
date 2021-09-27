#include <iostream>
using namespace std;
int main(){
int n;cin>>n;int v[105];for(int i=1;i<=n;i++){cin>>v[i];}
for(int i=n;i>=1;i--){cout<<v[i]<<" ";}
}