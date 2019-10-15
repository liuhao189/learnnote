# JS代码规范

良好的代码规范，可以使代码更加整洁，提高可读性，便于团队开发。

以下规范来自Airbnb。

# References

1、Use const for all of your references; avoid using var. eslint: prefer-const, no-const-assign。
确保不能修改你的引用，否则会导致bug或者难以理解的错误。

2、Reassign references, use let instead of var. eslint: no-var.
let块级作用域而不是像var的函数作用域。

# Objects

1、Use the literal syntax for object creation.eslint:no-new-object

2、Use computed property names when creating objects with dynamic property names.
允许你在一个地方定义所有的属性名。

3、Use object method shorthand. eslint:object-shorthand
使用对象方法的简写eslint:object-shorthand

4、Use property value shorthand.eslint:object-shorthand
可以使其写法和描述更短。

5、Group your shorthand properties at the begining of your object declaration.
更容易知道哪些属性是用了简写

6、Only quote properties that are invalid identifiers. eslint:quote-props.
更容易阅读，提升了语法的高亮显示，并且同样有利于JS引擎的优化。

7、Do not call object.prototype methods directly, such as hasOwnProperty, propertyIsEnumerable and isPrototypeOf.
eslint: no-prototype-builtins.
这些方法可能会被对象覆盖。eg{hasOwnProperty:false} or Object.crate(null)

```js
console.log(Object.prototype.hasOwnProperty.call(obj,key))
```

8、Perfer the object spread operator over Object.assign() to shallow-copy objects.Use thr object rest operator to get a new object with certain properties omitted.

# Arrays

1、Use the literal syntax for array creation. eslint: no-array-constructor.

2、Use Array#push instead of direct assignment to add items to an array.

3、Use array spreads ... to copy arrays.

4、To convert an iterable object to an array, use spreads ... instrad of Array.from.

5、Use Array.form for converting an array-like object to an array.

6、Use Array.from instead of spread ... for mapping over iterables, because it avoids creating an intermediate array.

7、Use return statements in array method callbacks. It is ok to omit the return if the function body consists of a single statement returning an expression without side effects eslint: array-callback-return.

8、Use line breaks after open and before close array brakets if an array has multiple lines.

```js
var romanToInt = function(s) {
    if(!s || !s.length) return;
    let specialNumMap={
        'IV':4,
        'IX':9,
        'XL':40,
        'XC':90,
        'CD':400,
        'CM':900
    }
    let numMap={
        'I':1,
        'V':5,
        'X':10,
        'L':50,
        'C':100,
        'D':500,
        'M':1000
    }
    let sum=0;
    for(let i=0;i<=s.length-1;++i){
        let currChar=s.charAt(i);
        if(i+1<s.length){
            let nextChar=s.charAt(i+1)   
            let specialStr=currChar+nextChar;
            if(specialNumMap[specialStr]){
                sum+=specialNumMap[specialStr];
                ++i;
            }else{
               sum+=    numMap[currChar];
            }
            continue;
        }
            sum+=numMap[currChar];
    }
    return sum;
};
```