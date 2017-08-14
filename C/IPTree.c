/*************************************************************************
	> Author: Juntaran
	> Mail: JuntaranMail@gmail.com
	> Created Time: 2017/7/29 20:01
 ************************************************************************/

#include <stdio.h>
#include <stdlib.h>

// 路由元素节点
typedef struct node {
	struct	node *left;
	struct	node *right;
	int		port;
} NODE;

// 创建节点
NODE *createNode() {
	NODE *pNode = malloc(sizeof(NODE));
	pNode -> left  = NULL;
	pNode -> right = NULL;
	pNode -> port = -1;
	return pNode;
}

// 路由表root元素
NODE *routeTree = NULL;

// 创建路由表
// 在这个路由表中，0向左，1向右
// 输入：路由IP 掩码位 端口号
void insertRoute(int iRoute, int iMask, int iPort) {
	int judge = 0;
	static int iInitFlag = 0;

	if (iInitFlag == 0) {
		routeTree = createNode();
		iInitFlag = 1;
	}

	printf("Input route: %8x, mask: %4d, port: %4d\n", iRoute, iMask, iPort);

	// 10.10.0.0	8		1993
	NODE *currentNode = routeTree;

	int i;
	for (i = 0; i < iMask; i++) {
		// 根据ip二进制从左到右按位解析
		judge = (iRoute >> (31 - i)) & 0x1;

		if(judge == 0) {
			if (NULL == currentNode->left) {
				currentNode->left = createNode();
			}
			currentNode = currentNode->left;
//			printf("0 left\n");
		} else {
			if (NULL == currentNode->right) {
				currentNode->right = createNode();
			}
			currentNode = currentNode->right;
//			printf("1 right\n");
		}
	}
	currentNode->port = iPort;
	printf("%d port\n", iPort);
	return;
}


// 定位路由节点
NODE* localteRoute(int ip) {
    int judge = 0, port = -1;
    NODE *currentNode = routeTree;
    port = (currentNode->port == -1) ? port : currentNode->port;

    printf("start locate ip: %8x\n", ip);

    int i = 0;
    for (i = 0; i < 32; i++) {
        judge = (ip >> (31-i)) & 0x1;

        if (judge == 0) {
            if (currentNode->left != NULL) {
                currentNode = currentNode->left;
                port = (currentNode->port == -1) ? port : currentNode->port;
//				printf("0 -> left, %d port\n", port);
            } else {
                break;
            }
        } else {
            if (NULL != currentNode->right) {
                currentNode = currentNode->right;
                port = (currentNode->port == -1) ? port : currentNode->port;
//				printf("1 -> right, %d port\n", port);
            } else {
                break;
            }
        }
    }
    return currentNode;
}


// 删除一条路由
void deleteRoute(int Route, int Mask) {
    int judge = 0;
    printf("Input route: %8x, mask: %4d\n", Route, Mask);

    NODE *currentNode = routeTree;
    int i;
    for (i = 0; i < Mask; i++) {
        // 根据ip二进制从左到右按位解析
        judge = (Route >> (31 - i)) & 0x1;

        if(judge == 0) {
            if (NULL == currentNode->left) {
                currentNode->left = createNode();
            }
            currentNode = currentNode->left;
        } else {
            if (NULL == currentNode->right) {
                currentNode->right = createNode();
            }
            currentNode = currentNode->right;
        }
    }
    // if (i < Mask) {
        // printf("Delete Error\n");
        // return;
    // }
    currentNode->port = -1;
    currentNode->left = NULL;
    currentNode->right = NULL;
    printf("Delete Success\n");
    return;
}


// 查找路由表函数
int searchRoute(int ip) {
    NODE* currentNode = localteRoute(ip);
    printf("Get %d\n", currentNode->port);
    return currentNode->port;
}

int main() {
	int a[4] = {0}, Route = 0,Mask = 0, Port = -1;

	// 构造路由表
	while(0 != scanf("%d.%d.%d.%d/%d %d\n", &a[0], &a[1], &a[2], &a[3], &Mask, &Port)) {
		printf("%d.%d.%d.%d/%d %d\n", a[0], a[1], a[2], a[3], Mask, Port);
		Route = (a[0] << 24) | (a[1] << 16) | (a[2] << 8) | (a[3] << 0);
        insertRoute(Route, Mask, Port);
	}

	// 输入ip查询port
    int ip[4] = {0};
    scanf("search: %d.%d.%d.%d", &ip[0], &ip[1], &ip[2], &ip[3]);
    printf("input ip: %d.%d.%d.%d\n", ip[0], ip[1], ip[2], ip[3]);
    int iIp = (ip[0] << 24) | (ip[1] << 16) | (ip[2] << 8) | (ip[3] << 0);
//	printf("%d\n", searchRoute(iIp));
    searchRoute(iIp);
    deleteRoute((a[0] << 24) | (a[1] << 16) | (a[2] << 8) | (a[3] << 0), Mask);
    searchRoute(iIp);
}