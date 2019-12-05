建造者模式属于创建型设计模式，它主要是将一个复杂对象的构建与表示分离,使用多个简单的对象一步一步构建成一个复杂的对象


Encapsulate creating and assembling the parts of a complex object in a separate Builder object.
A class delegates object creation to a Builder object instead of creating the objects directly.

A class (the same construction process) can delegate to different Builder objects to create different representations of a complex object.

# 定义
The intent of the Builder design pattern is to separate the construction of a complex object from its representation. By doing so the same construction process can create different representations.[1]

