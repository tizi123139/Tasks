# JavaScript

*这是一种轻量级、解释型、面向对象的脚本语言，主要被设计用于在网页上实现动态效果*

---



## 作用

- 客户端脚本：在用户浏览器中执行，实现动态效果和用户交互

- 网页开发：与HTML和css协同工作，使得网页有更强的交互性和动态性

- 后端开发：实现服务器端应用的开发

---



## 语法

- ### **常量**

##### 数字常量

3.14         1001

##### 字符串常量

可以使用单引号或双引号

"John Doe"

##### 表达式常量

5 + 6

##### 数组常量

[40, 100, 1, 5, 25, 10]

##### 函数常量

function myFunction(a, b) { return a * b;}



- ### **变量**

使用关键字 **var** 来定义变量， 使用等号来为变量赋值

var x, length

x = 5

length = 6



- ### **关键字**

| abstract | else       | instanceof | super        |
| -------- | ---------- | ---------- | ------------ |
| boolean  | enum       | int        | switch       |
| break    | export     | interface  | synchronized |
| byte     | extends    | let        | this         |
| case     | false      | long       | throw        |
| catch    | final      | native     | throws       |
| char     | finally    | new        | transient    |
| class    | float      | null       | true         |
| const    | for        | package    | try          |
| continue | function   | private    | typeof       |
| debugger | goto       | protected  | var          |
| default  | if         | public     | void         |
| delete   | implements | return     | volatile     |
| do       | import     | short      | while        |
| double   | in         | static     | with         |

- ### **注释**

双斜杠 **//** 后的内容将会被浏览器忽略



---



## 数据类型

**值类型(基本类型)**：字符串（String）、数字(Number)、布尔(Boolean)、空（Null）、未定义（Undefined）、Symbol。

**引用数据类型（对象类型）**：对象(Object)、数组(Array)、函数(Function)



JavaScript 拥有动态类型。这意味着相同的变量可用作不同的类型：

var x;        // x 为 undefined

var x = 5;      // 现在 x 为数字

var x = "John";   // 现在 x 为字符串



- ### 字符串

储存字符的变量

var carname="Volvo XC60";
var carname='Volvo XC60';



- ### 数字

var x1=34.00;   //使用小数点来写
var x2=34;     //不使用小数点来写



- ### 布尔

布尔（逻辑）只能有两个值：true 或 false。

var x=true;
var y=false;



- ### 数组

下面的代码创建名为 cars 的数组：

var cars=new Array();
cars[0]="Saab";
cars[1]="Volvo";
cars[2]="BMW";



- ### 对象

var person={firstname:"John", lastname:"Doe", id:5566};



- ### Undefined 和 Null

Undefined 这个值表示变量不含有值。

可以通过将变量的值设置为 null 来清空变量。