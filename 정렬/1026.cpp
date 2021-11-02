#include <stdio.h>
#include <vector>
#include <algorithm>

using namespace std;

bool comp(int a, int b) {
  return a > b;
}

int main() {
  int N, i, max=50, result=0;
  int a[max], b[max];
  scanf("%d", &N);
  for (i=0; i<N; i++) scanf("%d", &a[i]);
  for (i=0; i<N; i++) scanf("%d", &b[i]);
  sort(a, a+N);
  sort(b, b+N, comp);

  for (i=0; i<N; i++) result += a[i] * b[i];
  printf("%d", result);
}
