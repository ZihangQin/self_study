package com.company;

import java.util.Scanner;

public class Main {
    public static void main(String[] args){
/*        Scanner scanner = new Scanner(System.in);
        int length = scanner.nextInt();*/

        //声明并初始化二维数组
        int[] [] yh = new int[10][];

        //赋值
        for (int i = 0; i <yh.length;i++ ){
            yh[i] = new int[i+1];
            yh[i] [0]=1;
            yh[i] [i]=1;
            if(i>1){
                for(int j=1;j<i;j++){
                    yh[i] [j] = yh[i-1][j-1]+yh[i-1][j];
                }

            }        }

        //遍历
        for(int i=0;i<yh.length;i++){
            for (int j = 0; j <yh[i].length;j++ ){
                System.out.print(yh[i][j]+ "   ");
            }
            System.out.println("  ");
        }
    }
}
