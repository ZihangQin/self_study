package deno.cn.qianfianfeng;

import java.util.Scanner;

public class FileEncry {
    public static void main(String[] args) {
        //提示用户输入铭文密码
        System.out.println("请输入铭文密码：");
        //Scanner类表示一个文本扫描器，他可以扫描从键盘上输入的字符
        Scanner in = new Scanner(System.in);
        //方法nextLine（）返回键盘上输入的一行字符串
        String secreterStr = in.nextLine();
        //将字符串转化为数组
        char[] secreterChar = secreterStr.toCharArray();
        //字符变量用于保存加密密钥
        char secret = 'x';
        //加密运算，将要加密的字符与x进行按位异或运算得到密文
        System.out.print("密文：");
        for (int i = 0; i < secreterChar.length; i++) {
         // secreterChar[i]表示字符串中的数组
            secreterChar[i] = (char) (secreterChar[i] ^ secret);
            System.out.print(secreterChar[i]);
        }
        //解密运算，与加密运算逻辑相反
        System.out.print("\n 铭文 ：");
        for (int j = 0; j < secreterChar.length; j++) {
            secreterChar[j] = (char) (secreterChar[j] ^ secret);
            System.out.print(secreterChar[j]);
        }

    }
}
