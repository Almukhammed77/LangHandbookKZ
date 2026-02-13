package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Almukhammed77/LangHandbookKZ/models"
	"github.com/Almukhammed77/LangHandbookKZ/storage"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

const jwtSecret = "my-secret-key-12345"

type TutorialSection struct {
	Title   string
	Slug    string
	Content string
}

// JavaScript Tutorials - 8 модулей
var jsTutorials = []TutorialSection{
	{
		Title: "JavaScript Home",
		Slug:  "js-home",
		Content: `<h2>Добро пожаловать в JavaScript</h2>
		<p>JavaScript — это язык программирования, который делает веб-страницы интерактивными. Создан в 1995 году Бренданом Эйхом.</p>
		
		<h3>Что можно делать с JavaScript?</h3>
		<ul>
			<li>Добавлять интерактивность на сайты</li>
			<li>Создавать веб-приложения (React, Vue, Angular)</li>
			<li>Писать серверный код (Node.js)</li>
			<li>Создавать мобильные приложения (React Native)</li>
			<li>Разрабатывать игры</li>
		</ul>
		
		<div style="background: #f0f9ff; padding: 20px; border-radius: 12px; margin: 20px 0;">
			<h3>JavaScript в цифрах:</h3>
			<p>✅ 98% всех сайтов используют JavaScript</p>
			<p>✅ Более 15 миллионов разработчиков</p>
			<p>✅ 1.5+ миллиона пакетов в npm</p>
		</div>`,
	},
	{
		Title: "JS Introduction",
		Slug:  "js-introduction",
		Content: `<h2>Введение в JavaScript</h2>
		<p>JavaScript (JS) — это высокоуровневый, интерпретируемый язык программирования.</p>
		
		<h3>Где используется JavaScript?</h3>
		<ul>
			<li><strong>Frontend:</strong> React, Vue, Angular, Svelte</li>
			<li><strong>Backend:</strong> Node.js, Deno, Bun</li>
			<li><strong>Мобильные приложения:</strong> React Native, Ionic</li>
			<li><strong>Десктоп:</strong> Electron (VS Code, Slack, Discord)</li>
		</ul>
		
		<div style="background: #fffbeb; border-left: 6px solid #f59e0b; padding: 20px; margin: 20px 0;">
			<p><strong>Интересный факт:</strong> JavaScript был создан за 10 дней!</p>
		</div>`,
	},
	{
		Title: "JS Syntax",
		Slug:  "js-syntax",
		Content: `<h2>Синтаксис JavaScript</h2>
		
		<h3>Первая программа</h3>
		<pre><code>console.log("Hello, World!");
alert("Hello, World!");</code></pre>
		
		<h3>Переменные</h3>
		<pre><code>// let - переменная (можно изменять)
let age = 25;
age = 26; // ✅ можно

// const - константа (нельзя изменять)
const birthYear = 1999;

// var - старый способ (не используйте)
var oldWay = "не рекомендуется";</code></pre>
		
		<h3>Типы данных</h3>
		<pre><code>let name = "Алматы";        // String
let count = 42;            // Number
let isActive = true;       // Boolean
let user = { name: "Азамат" }; // Object
let languages = ["JS", "Python"]; // Array</code></pre>`,
	},
	{
		Title: "JS Functions",
		Slug:  "js-functions",
		Content: `<h2>Функции в JavaScript</h2>
		
		<h3>Function Declaration</h3>
		<pre><code>function greet(name) {
    return "Привет, " + name + "!";
}</code></pre>
		
		<h3>Arrow Functions (ES6+)</h3>
		<pre><code>const greet = name => "Привет, " + name + "!";
const sum = (a, b) => a + b;
const sayHello = () => "Сәлем!";</code></pre>
		
		<h3>Параметры по умолчанию</h3>
		<pre><code>function greet(name = "гость") {
    return "Добро пожаловать, " + name;
}</code></pre>`,
	},
	{
		Title: "JS Arrays",
		Slug:  "js-arrays",
		Content: `<h2>Массивы в JavaScript</h2>
		
		<h3>Создание массивов</h3>
		<pre><code>let fruits = ["Яблоко", "Банан", "Апельсин"];</code></pre>
		
		<h3>Основные методы</h3>
		<pre><code>fruits.push("Груша");      // Добавить в конец
fruits.pop();             // Удалить с конца
fruits.indexOf("Банан");  // Индекс элемента</code></pre>
		
		<h3>🔥 Современные методы (ES6+)</h3>
		<pre><code>const numbers = [1, 2, 3, 4, 5];

// map - преобразует каждый элемент
const doubled = numbers.map(n => n * 2);
// [2, 4, 6, 8, 10]

// filter - фильтрует элементы
const evens = numbers.filter(n => n % 2 === 0);
// [2, 4]</code></pre>`,
	},
	{
		Title: "JS Objects",
		Slug:  "js-objects",
		Content: `<h2>Объекты в JavaScript</h2>
		
		<h3>Создание объектов</h3>
		<pre><code>const user = {
    name: "Диас",
    age: 25,
    city: "Астана",
    greet() {
        return "Сәлем, меня зовут " + this.name;
    }
};</code></pre>
		
		<h3>Деструктуризация (ES6)</h3>
		<pre><code>const { name, age } = user;
console.log(name); // "Диас"</code></pre>`,
	},
	{
		Title: "JS DOM",
		Slug:  "js-dom",
		Content: `<h2>DOM — работа со страницей</h2>
		
		<h3>Поиск элементов</h3>
		<pre><code>document.getElementById("header");
document.querySelector(".my-class");
document.querySelectorAll("div.item");</code></pre>
		
		<h3>Изменение содержимого</h3>
		<pre><code>element.textContent = "Новый текст";
element.innerHTML = "&lt;strong&gt;Жирный текст&lt;/strong&gt;";
element.style.color = "red";</code></pre>
		
		<h3>События</h3>
		<pre><code>button.addEventListener("click", function() {
    console.log("Клик!");
});</code></pre>`,
	},
	{
		Title: "JS Async",
		Slug:  "js-async",
		Content: `<h2>Асинхронность в JavaScript</h2>
		
		<h3>Promise (ES6)</h3>
		<pre><code>fetch("https://api.example.com")
    .then(res => res.json())
    .then(data => console.log(data));</code></pre>
		
		<h3>Async/Await (ES2017)</h3>
		<pre><code>async function getData() {
    const res = await fetch("https://api.example.com");
    const data = await res.json();
    console.log(data);
}</code></pre>`,
	},
}

// Go Tutorials - 9 модулей
var goTutorials = []TutorialSection{
	{
		Title: "Go Home",
		Slug:  "go-home",
		Content: `<h2>Добро пожаловать в Go!</h2>
		<p>Go (Golang) — современный язык программирования от Google. Создан в 2009 году Робертом Гриземером, Робом Пайком и Кеном Томпсоном.</p>
		
		<div style="background: #e0f2fe; padding: 20px; border-radius: 12px; margin: 20px 0;">
			<h3>Почему Go?</h3>
			<ul>
				<li>Молниеносная компиляция</li>
				<li>Легковесные горутины</li>
				<li>Единый бинарник без зависимостей</li>
				<li>Безопасность памяти</li>
				<li>Используют: Google, Uber, Twitch, Dropbox</li>
			</ul>
		</div>
		
		<p>В Казахстане Go активно используют Kaspi.kz, Chocofamily, Halyk Bank.</p>`,
	},
	{
		Title: "Go Introduction",
		Slug:  "go-introduction",
		Content: `<h2>Введение в Go</h2>
		
		<h3>Философия Go</h3>
		<p>Go создавался как ответ на сложность современных языков. Основные принципы:</p>
		<ul>
			<li><strong>Простота:</strong> всего 25 ключевых слов!</li>
			<li><strong>Читаемость:</strong> единый стиль кода (gofmt форматирует код автоматически)</li>
			<li><strong>Скорость:</strong> быстрая компиляция и выполнение</li>
			<li><strong>Параллелизм:</strong> горутины вместо потоков</li>
		</ul>
		
		<h3>Где используется Go?</h3>
		<ul>
			<li>Веб-серверы и API</li>
			<li>Микросервисы</li>
			<li>DevOps инструменты (Docker, Kubernetes)</li>
			<li>Базы данных</li>
			<li>Блокчейн</li>
		</ul>`,
	},
	{
		Title: "Go Get Started",
		Slug:  "go-get-started",
		Content: `<h2>Начало работы с Go</h2>
		
		<h3>Установка Go</h3>
		<ol>
			<li>Перейди на <a href="https://go.dev/dl/">golang.org/dl</a></li>
			<li>Скачай версию для своей ОС</li>
			<li>Установи и проверь: <code>go version</code></li>
		</ol>
		
		<h3>Первая программа</h3>
		<pre><code>package main

import "fmt"

func main() {
    fmt.Println("Сәлем, Go!")
}</code></pre>
		
		<h3>Запуск программы</h3>
		<pre><code>go run main.go</code></pre>`,
	},
	{
		Title: "Go Syntax",
		Slug:  "go-syntax",
		Content: `<h2>Синтаксис Go</h2>
		
		<h3>Основы</h3>
		<pre><code>package main

import "fmt"

func main() {
    var name string = "Алматы"
    var age int = 170
    city := "Астана"
    
    fmt.Printf("Привет, %s! Возраст: %d\n", name, age)
}</code></pre>
		
		<h3>Типы данных</h3>
		<pre><code>var i int = 42
var i8 int8 = 127
var u uint = 42
var f32 float32 = 3.14
var f64 float64 = 3.14159
var s string = "Текст"
var b bool = true</code></pre>`,
	},
	{
		Title: "Go Variables",
		Slug:  "go-variables",
		Content: `<h2>Переменные в Go</h2>
		
		<h3>Объявление переменных</h3>
		<pre><code>var name string = "Азамат"
var city = "Алматы"
age := 25</code></pre>
		
		<h3>Константы</h3>
		<pre><code>const Pi = 3.14159
const StatusOK = 200</code></pre>`,
	},
	{
		Title: "Go Functions",
		Slug:  "go-functions",
		Content: `<h2>Функции в Go</h2>
		
		<h3>Объявление функций</h3>
		<pre><code>func greet(name string) string {
    return "Привет, " + name + "!"
}</code></pre>
		
		<h3>Множественные возвращаемые значения</h3>
		<pre><code>func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("деление на ноль")
    }
    return a / b, nil
}</code></pre>`,
	},
	{
		Title: "Go Arrays",
		Slug:  "go-arrays",
		Content: `<h2>Массивы и Срезы</h2>
		
		<h3>Массивы</h3>
		<pre><code>var arr [3]int = [3]int{1, 2, 3}</code></pre>
		
		<h3>Срезы (Slices)</h3>
		<pre><code>slice := []int{1, 2, 3}
slice = append(slice, 4, 5)</code></pre>
		
		<h3>Итерация</h3>
		<pre><code>for i, fruit := range fruits {
    fmt.Println(i, fruit)
}</code></pre>`,
	},
	{
		Title: "Go Maps",
		Slug:  "go-maps",
		Content: `<h2>Карты (Maps) в Go</h2>
		
		<h3>Создание map</h3>
		<pre><code>user := map[string]string{
    "name": "Азамат",
    "city": "Алматы",
}</code></pre>
		
		<h3>Операции с map</h3>
		<pre><code>name := user["name"]
delete(user, "job")</code></pre>`,
	},
	{
		Title: "Go Structs",
		Slug:  "go-structs",
		Content: `<h2>Структуры в Go</h2>
		
		<h3>Определение структуры</h3>
		<pre><code>type User struct {
    ID   int
    Name string
    Age  int
}</code></pre>
		
		<h3>Создание экземпляров</h3>
		<pre><code>user1 := User{ID: 1, Name: "Айгерим", Age: 25}</code></pre>
		
		<h3>Методы структур</h3>
		<pre><code>func (u User) Greet() string {
    return "Привет, я " + u.Name
}</code></pre>`,
	},
}

// TypeScript Tutorials - 9 модулей
var tsTutorials = []TutorialSection{
	{
		Title: "TypeScript Home",
		Slug:  "ts-home",
		Content: `<h2>Добро пожаловать в TypeScript!</h2>
		<p>TypeScript — это строгая типизация для JavaScript. Создан Microsoft в 2012 году.</p>
		
		<div style="background: #e6f0ff; padding: 20px; border-radius: 12px; margin: 20px 0;">
			<h3>Зачем TypeScript?</h3>
			<ul>
				<li>Находит ошибки до выполнения кода</li>
				<li>Автодополнение в IDE</li>
				<li>Рефакторинг без страха</li>
			</ul>
		</div>`,
	},
	{
		Title: "TS Introduction",
		Slug:  "ts-introduction",
		Content: `<h2>Введение в TypeScript</h2>
		
		<h3>Что такое TypeScript?</h3>
		<p>TypeScript добавляет в JavaScript статическую типизацию, классы, интерфейсы, дженерики.</p>
		
		<div style="background: #eff6ff; border-left: 6px solid #2563eb; padding: 20px; margin: 20px 0;">
			<p><strong>TypeScript vs JavaScript:</strong></p>
			<pre><code>// JavaScript
function add(a, b) { return a + b; }
add(5, "10");  // "510" 

// TypeScript
function add(a: number, b: number): number {
    return a + b;
}</code></pre>
		</div>`,
	},
	{
		Title: "TS Basic Types",
		Slug:  "ts-basic-types",
		Content: `<h2>Базовые типы TypeScript</h2>
		
		<h3>Примитивные типы</h3>
		<pre><code>let name: string = "Алматы";
let age: number = 28;
let isActive: boolean = true;
let numbers: number[] = [1, 2, 3];</code></pre>`,
	},
	{
		Title: "TS Functions",
		Slug:  "ts-functions",
		Content: `<h2>Функции в TypeScript</h2>
		
		<h3>Типизация параметров</h3>
		<pre><code>function greet(name: string): string {
    return "Привет, " + name;
}</code></pre>
		
		<h3>Опциональные параметры</h3>
		<pre><code>function createUser(name: string, age?: number): string {
    if (age) {
        return "Имя: " + name + ", Возраст: " + age;
    }
    return "Имя: " + name;
}</code></pre>`,
	},
	{
		Title: "TS Interfaces",
		Slug:  "ts-interfaces",
		Content: `<h2>Интерфейсы в TypeScript</h2>
		
		<h3>Базовый интерфейс</h3>
		<pre><code>interface User {
    id: number;
    name: string;
    email?: string;
}

const user: User = {
    id: 1,
    name: "Айжан"
};</code></pre>`,
	},
	{
		Title: "TS Classes",
		Slug:  "ts-classes",
		Content: `<h2>Классы в TypeScript</h2>
		
		<h3>Модификаторы доступа</h3>
		<pre><code>class Person {
    public name: string;
    private age: number;
    
    constructor(name: string, age: number) {
        this.name = name;
        this.age = age;
    }
    
    public greet(): string {
        return "Привет, я " + this.name;
    }
}</code></pre>`,
	},
	{
		Title: "TS Generics",
		Slug:  "ts-generics",
		Content: `<h2>Дженерики в TypeScript</h2>
		
		<h3>Generic функции</h3>
		<pre><code>function identity&lt;T&gt;(arg: T): T {
    return arg;
}

identity&lt;string&gt;("text");
identity&lt;number&gt;(42);</code></pre>`,
	},
	{
		Title: "TS Modules",
		Slug:  "ts-modules",
		Content: `<h2>Модули в TypeScript</h2>
		
		<h3>Экспорт и импорт</h3>
		<pre><code>// math.ts
export function add(a: number, b: number): number {
    return a + b;
}

// app.ts
import { add } from "./math.js";</code></pre>`,
	},
	{
		Title: "TS Utility Types",
		Slug:  "ts-utility",
		Content: `<h2>Utility Types</h2>
		
		<pre><code>interface User {
    id: number;
    name: string;
    email: string;
}

// Partial — все поля опциональны
type PartialUser = Partial&lt;User&gt;;</code></pre>`,
	},
}

// Python Tutorials - 7 модулей
var pythonTutorials = []TutorialSection{
	{
		Title: "Python Home",
		Slug:  "python-home",
		Content: `<h2>Добро пожаловать в Python</h2>
		<p>Python — самый популярный язык для начинающих. Создан Гвидо ван Россумом в 1991 году.</p>
		
		<div style="background: #e6f7e6; padding: 20px; border-radius: 12px; margin: 20px 0;">
			<h3>Почему Python?</h3>
			<ul>
				<li>Простой и понятный синтаксис</li>
				<li>Огромное количество библиотек</li>
				<li>Используется в AI, Data Science, Web</li>
			</ul>
		</div>`,
	},
	{
		Title: "Python Introduction",
		Slug:  "python-introduction",
		Content: `<h2>Введение в Python</h2>
		
		<h3>Где используется Python?</h3>
		<ul>
			<li>Веб-разработка (Django, Flask)</li>
			<li>Машинное обучение (TensorFlow)</li>
			<li>Анализ данных (Pandas)</li>
			<li>Автоматизация</li>
		</ul>`,
	},
	{
		Title: "Python Syntax",
		Slug:  "python-syntax",
		Content: `<h2>Python Syntax</h2>
		
		<h3>Первая программа</h3>
		<pre><code>print("Hello, World!")</code></pre>
		
		<h3>Python Indentation</h3>
		<p><strong>Python uses indentation to indicate a block of code.</strong></p>
		
		<pre><code>if 5 > 2:
    print("Five is greater than two!")</code></pre>`,
	},
	{
		Title: "Python Variables",
		Slug:  "python-variables",
		Content: `<h2>Python Variables</h2>
		
		<h3>Creating Variables</h3>
		<pre><code>x = 5
y = "John"
print(x)
print(y)</code></pre>`,
	},
	{
		Title: "Python Data Types",
		Slug:  "python-datatypes",
		Content: `<h2>Python Data Types</h2>
		
		<h3>Built-in Data Types</h3>
		<ul>
			<li>str - текст</li>
			<li>int, float - числа</li>
			<li>list - список</li>
			<li>dict - словарь</li>
			<li>bool - True/False</li>
		</ul>`,
	},
	{
		Title: "Python If...Else",
		Slug:  "python-ifelse",
		Content: `<h2>Python If...Else</h2>
		
		<h3>If statement:</h3>
		<pre><code>a = 33
b = 200
if b > a:
    print("b is greater than a")</code></pre>`,
	},
	{
		Title: "Python Loops",
		Slug:  "python-loops",
		Content: `<h2>Python Loops</h2>
		
		<h3>For Loops</h3>
		<pre><code>fruits = ["apple", "banana", "cherry"]
for x in fruits:
    print(x)</code></pre>`,
	},
}

// Java Tutorials - 8 модулей
var javaTutorials = []TutorialSection{
	{
		Title: "Java Home",
		Slug:  "java-home",
		Content: `<h2>Добро пожаловать в Java!</h2>
		<p>Java — объектно-ориентированный язык программирования, созданный Sun Microsystems в 1995 году.</p>
		
		<div style="background: #fef3c7; padding: 20px; border-radius: 12px; margin: 20px 0;">
			<h3>Java в цифрах:</h3>
			<ul>
				<li>Более 10 миллионов разработчиков</li>
				<li>Работает на 3 миллиардах устройств</li>
				<li>Android-разработка базируется на Java</li>
			</ul>
		</div>`,
	},
	{
		Title: "Java Introduction",
		Slug:  "java-introduction",
		Content: `<h2>Введение в Java</h2>
		
		<h3>Философия Java</h3>
		<p>Основные принципы Java:</p>
		<ul>
			<li><strong>WORA</strong> — Write Once, Run Anywhere</li>
			<li><strong>Объектно-ориентированный</strong></li>
			<li><strong>Автоматическое управление памятью</strong></li>
		</ul>`,
	},
	{
		Title: "Java Syntax",
		Slug:  "java-syntax",
		Content: `<h2>Синтаксис Java</h2>
		
		<h3>Первая программа</h3>
		<pre><code>public class Main {
    public static void main(String[] args) {
        System.out.println("Сәлем, Java!");
    }
}</code></pre>`,
	},
	{
		Title: "Java Variables",
		Slug:  "java-variables",
		Content: `<h2>Переменные в Java</h2>
		
		<h3>Объявление переменных</h3>
		<pre><code>int age = 25;
double price = 99.99;
String name = "Азамат";
boolean isActive = true;</code></pre>`,
	},
	{
		Title: "Java Operators",
		Slug:  "java-operators",
		Content: `<h2>Операторы в Java</h2>
		
		<h3>Арифметические операторы</h3>
		<pre><code>int a = 10, b = 3;
int sum = a + b;      // 13
int diff = a - b;     // 7
int mult = a * b;     // 30
int div = a / b;      // 3</code></pre>`,
	},
	{
		Title: "Java Control Flow",
		Slug:  "java-control-flow",
		Content: `<h2>Управляющие конструкции</h2>
		
		<h3>If-else</h3>
		<pre><code>if (age >= 18) {
    System.out.println("Взрослый");
} else {
    System.out.println("Ребенок");
}</code></pre>`,
	},
	{
		Title: "Java Loops",
		Slug:  "java-loops",
		Content: `<h2>Циклы в Java</h2>
		
		<h3>For loop</h3>
		<pre><code>for (int i = 0; i < 5; i++) {
    System.out.println(i);
}</code></pre>`,
	},
	{
		Title: "Java Arrays",
		Slug:  "java-arrays",
		Content: `<h2>Массивы в Java</h2>
		
		<h3>Объявление массивов</h3>
		<pre><code>int[] numbers = new int[5];
int[] arr = {1, 2, 3, 4, 5};</code></pre>`,
	},
}

// C# Tutorials - 7 модулей
var csharpTutorials = []TutorialSection{
	{
		Title: "C# Home",
		Slug:  "csharp-home",
		Content: `<h2>Добро пожаловать в C#!</h2>
		<p>C# — современный объектно-ориентированный язык от Microsoft, созданный в 2000 году.</p>
		
		<h3>Где используется C#?</h3>
		<ul>
			<li>Разработка игр (Unity)</li>
			<li>Веб-приложения (ASP.NET)</li>
			<li>Десктопные приложения</li>
		</ul>`,
	},
	{
		Title: "C# Introduction",
		Slug:  "csharp-introduction",
		Content: `<h2>Введение в C#</h2>
		
		<h3>Первая программа</h3>
		<pre><code>using System;

class Program {
    static void Main() {
        Console.WriteLine("Сәлем, C#!");
    }
}</code></pre>`,
	},
	{
		Title: "C# Syntax",
		Slug:  "csharp-syntax",
		Content: `<h2>Синтаксис C#</h2>
		
		<h3>Переменные</h3>
		<pre><code>int age = 25;
string name = "Азамат";
bool isActive = true;</code></pre>`,
	},
	{
		Title: "C# Operators",
		Slug:  "csharp-operators",
		Content: `<h2>Операторы в C#</h2>
		
		<pre><code>int a = 10, b = 3;
int sum = a + b;
int diff = a - b;</code></pre>`,
	},
	{
		Title: "C# Control Flow",
		Slug:  "csharp-control-flow",
		Content: `<h2>Управляющие конструкции</h2>
		
		<pre><code>if (age >= 18) {
    Console.WriteLine("Взрослый");
}</code></pre>`,
	},
	{
		Title: "C# Loops",
		Slug:  "csharp-loops",
		Content: `<h2>Циклы в C#</h2>
		
		<pre><code>for (int i = 0; i < 5; i++) {
    Console.WriteLine(i);
}</code></pre>`,
	},
	{
		Title: "C# Arrays",
		Slug:  "csharp-arrays",
		Content: `<h2>Массивы в C#</h2>
		
		<pre><code>int[] numbers = new int[5];
int[] arr = {1, 2, 3, 4, 5};</code></pre>`,
	},
}

// Rust Tutorials - 7 модулей
var rustTutorials = []TutorialSection{
	{
		Title: "Rust Home",
		Slug:  "rust-home",
		Content: `<h2>Добро пожаловать в Rust!</h2>
		<p>Rust — системный язык программирования от Mozilla, созданный в 2010 году.</p>
		
		<h3>Особенности Rust</h3>
		<ul>
			<li>Безопасность памяти</li>
			<li>Высокая производительность</li>
			<li>Современный синтаксис</li>
		</ul>`,
	},
	{
		Title: "Rust Introduction",
		Slug:  "rust-introduction",
		Content: `<h2>Введение в Rust</h2>
		
		<h3>Первая программа</h3>
		<pre><code>fn main() {
    println!("Сәлем, Rust!");
}</code></pre>`,
	},
	{
		Title: "Rust Variables",
		Slug:  "rust-variables",
		Content: `<h2>Переменные в Rust</h2>
		
		<pre><code>let x = 5;        // неизменяемая
let mut y = 5;    // изменяемая
y = 6;</code></pre>`,
	},
	{
		Title: "Rust Functions",
		Slug:  "rust-functions",
		Content: `<h2>Функции в Rust</h2>
		
		<pre><code>fn add(x: i32, y: i32) -> i32 {
    x + y
}</code></pre>`,
	},
	{
		Title: "Rust Ownership",
		Slug:  "rust-ownership",
		Content: `<h2>Владение (Ownership)</h2>
		
		<pre><code>let s1 = String::from("hello");
let s2 = s1;  // владение перемещается</code></pre>`,
	},
	{
		Title: "Rust Structs",
		Slug:  "rust-structs",
		Content: `<h2>Структуры в Rust</h2>
		
		<pre><code>struct User {
    name: String,
    age: u32,
}</code></pre>`,
	},
	{
		Title: "Rust Enums",
		Slug:  "rust-enums",
		Content: `<h2>Перечисления (Enums)</h2>
		
		<pre><code>enum Direction {
    Up,
    Down,
    Left,
    Right,
}</code></pre>`,
	},
}

// C++ Tutorials - 7 модулей
var cppTutorials = []TutorialSection{
	{
		Title: "C++ Home",
		Slug:  "cpp-home",
		Content: `<h2>Добро пожаловать в C++!</h2>
		<p>C++ — язык программирования, созданный Бьёрном Страуструпом в 1985 году.</p>
		
		<h3>Где используется C++?</h3>
		<ul>
			<li>Игровые движки</li>
			<li>Браузеры</li>
			<li>Операционные системы</li>
		</ul>`,
	},
	{
		Title: "C++ Introduction",
		Slug:  "cpp-introduction",
		Content: `<h2>Введение в C++</h2>
		
		<h3>Первая программа</h3>
		<pre><code>#include <iostream>

int main() {
    std::cout << "Сәлем, C++!" << std::endl;
    return 0;
}</code></pre>`,
	},
	{
		Title: "C++ Syntax",
		Slug:  "cpp-syntax",
		Content: `<h2>Синтаксис C++</h2>
		
		<pre><code>int age = 25;
double price = 99.99;
std::string name = "Азамат";</code></pre>`,
	},
	{
		Title: "C++ Functions",
		Slug:  "cpp-functions",
		Content: `<h2>Функции в C++</h2>
		
		<pre><code>int add(int a, int b) {
    return a + b;
}</code></pre>`,
	},
	{
		Title: "C++ Classes",
		Slug:  "cpp-classes",
		Content: `<h2>Классы в C++</h2>
		
		<pre><code>class Person {
public:
    std::string name;
    int age;
};</code></pre>`,
	},
	{
		Title: "C++ Pointers",
		Slug:  "cpp-pointers",
		Content: `<h2>Указатели в C++</h2>
		
		<pre><code>int x = 42;
int* ptr = &x;
*ptr = 100;</code></pre>`,
	},
	{
		Title: "C++ Vectors",
		Slug:  "cpp-vectors",
		Content: `<h2>Векторы в C++</h2>
		
		<pre><code>#include <vector>
std::vector<int> v = {1, 2, 3};
v.push_back(4);</code></pre>`,
	},
}

// Swift Tutorials - 7 модулей
var swiftTutorials = []TutorialSection{
	{
		Title: "Swift Home",
		Slug:  "swift-home",
		Content: `<h2>Добро пожаловать в Swift!</h2>
		<p>Swift — современный язык программирования от Apple, представленный в 2014 году.</p>
		
		<h3>Где используется Swift?</h3>
		<ul>
			<li>iOS приложения</li>
			<li>macOS приложения</li>
			<li>watchOS приложения</li>
		</ul>`,
	},
	{
		Title: "Swift Introduction",
		Slug:  "swift-introduction",
		Content: `<h2>Введение в Swift</h2>
		
		<h3>Первая программа</h3>
		<pre><code>print("Сәлем, Swift!")</code></pre>`,
	},
	{
		Title: "Swift Variables",
		Slug:  "swift-variables",
		Content: `<h2>Переменные в Swift</h2>
		
		<pre><code>var age = 25        // изменяемая
let name = "Азамат"  // константа</code></pre>`,
	},
	{
		Title: "Swift Optionals",
		Slug:  "swift-optionals",
		Content: `<h2>Опциональные типы</h2>
		
		<pre><code>var age: Int? = 25
if let age = age {
    print(age)
}</code></pre>`,
	},
	{
		Title: "Swift Functions",
		Slug:  "swift-functions",
		Content: `<h2>Функции в Swift</h2>
		
		<pre><code>func greet(name: String) -> String {
    return "Привет, \(name)!"
}</code></pre>`,
	},
	{
		Title: "Swift Classes",
		Slug:  "swift-classes",
		Content: `<h2>Классы в Swift</h2>
		
		<pre><code>class Person {
    var name: String
    var age: Int
    
    init(name: String, age: Int) {
        self.name = name
        self.age = age
    }
}</code></pre>`,
	},
	{
		Title: "Swift Structs",
		Slug:  "swift-structs",
		Content: `<h2>Структуры в Swift</h2>
		
		<pre><code>struct Point {
    var x: Double
    var y: Double
}</code></pre>`,
	},
}

// Объединяем все туториалы
var tutorials = map[string][]TutorialSection{
	"Go":         goTutorials,
	"JavaScript": jsTutorials,
	"TypeScript": tsTutorials,
	"Python":     pythonTutorials,
	"Java":       javaTutorials,
	"C#":         csharpTutorials,
	"Rust":       rustTutorials,
	"C++":        cppTutorials,
	"Swift":      swiftTutorials,
}

func main() {
	storage.InitDB()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("templates/static"))))

	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/search", SearchHandler)
	http.HandleFunc("/register", RegisterHandler)
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/logout", LogoutHandler)
	http.HandleFunc("/language/", LanguageDetailHandler)
	http.HandleFunc("/profile", ProfileHandler)
	http.HandleFunc("/profile/update", ProfileUpdateHandler)
	http.HandleFunc("/profile/change-password", ProfileChangePasswordHandler)

	log.Println("Сервер запущен → http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func getUsernameFromCookie(r *http.Request) string {
	cookie, err := r.Cookie("auth_token")
	if err != nil {
		return ""
	}

	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil || !token.Valid {
		return ""
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if username, ok := claims["username"].(string); ok {
			return username
		}
	}
	return ""
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	username := getUsernameFromCookie(r)
	languages := storage.GetAllLanguages("", "popularity DESC")

	data := struct {
		Username  string
		Languages []*models.Language
		Query     string
	}{
		Username:  username,
		Languages: languages,
		Query:     "",
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Ошибка шаблона: "+err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
	tmpl.Execute(w, data)
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	username := getUsernameFromCookie(r)
	languages := storage.SearchLanguages(query)

	data := struct {
		Username  string
		Languages []*models.Language
		Query     string
	}{
		Username:  username,
		Languages: languages,
		Query:     query,
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Ошибка шаблона", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}

func LanguageDetailHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/language/")
	parts := strings.Split(path, "/")

	if len(parts) < 1 || parts[0] == "" {
		http.NotFound(w, r)
		return
	}

	idStr := parts[0]
	sectionSlug := "go-home"
	if len(parts) > 1 && parts[1] != "" {
		sectionSlug = parts[1]
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	lang := storage.GetLanguageByID(uint(id))
	if lang == nil {
		http.Error(w, "Язык не найден", http.StatusNotFound)
		return
	}

	views := lang.Views + 1
	storage.UpdateViews(uint(id), views)
	lang.Views = views

	username := getUsernameFromCookie(r)

	langTutorial, ok := tutorials[lang.Name]
	if !ok {
		langTutorial = []TutorialSection{
			{Title: "Главная", Slug: "home", Content: "<h2>Туториал в разработке</h2><p>Для языка " + lang.Name + " контент скоро появится.</p>"},
		}
	}

	currentContent := ""
	currentTitle := ""
	found := false

	for _, sec := range langTutorial {
		if sec.Slug == sectionSlug {
			currentContent = sec.Content
			currentTitle = sec.Title
			found = true
			break
		}
	}

	if !found && len(langTutorial) > 0 {
		currentContent = langTutorial[0].Content
		currentTitle = langTutorial[0].Title
		sectionSlug = langTutorial[0].Slug
	}

	data := struct {
		Username     string
		Language     *models.Language
		Sections     []TutorialSection
		SectionTitle string
		SectionSlug  string
		Content      template.HTML
		Query        string
	}{
		Username:     username,
		Language:     lang,
		Sections:     langTutorial,
		SectionTitle: currentTitle,
		SectionSlug:  sectionSlug,
		Content:      template.HTML(currentContent),
		Query:        "",
	}

	tmpl, err := template.ParseFiles("templates/language.html")
	if err != nil {
		http.Error(w, "Ошибка шаблона: "+err.Error(), http.StatusInternalServerError)
		log.Println("Template parse error:", err)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Ошибка рендеринга: "+err.Error(), http.StatusInternalServerError)
		log.Println("Execute error:", err)
		return
	}
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/register.html"))

	if r.Method == "GET" {
		tmpl.Execute(w, nil)
		return
	}

	r.ParseForm()
	username := strings.TrimSpace(r.FormValue("username"))
	password := strings.TrimSpace(r.FormValue("password"))

	if username == "" || password == "" {
		tmpl.Execute(w, map[string]string{"Error": "Заполните все поля"})
		return
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		tmpl.Execute(w, map[string]string{"Error": "Ошибка сервера"})
		return
	}

	user := models.User{
		Username:   username,
		Password:   string(hashed),
		Email:      username + "@example.com",
		FullName:   username,
		Role:       "user",
		Level:      1,
		Experience: 0,
		CreatedAt:  time.Now(),
	}

	if err := storage.DB.Create(&user).Error; err != nil {
		tmpl.Execute(w, map[string]string{"Error": "Такой логин уже занят"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	})
	tokenStr, _ := token.SignedString([]byte(jwtSecret))

	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    tokenStr,
		Path:     "/",
		MaxAge:   86400,
		HttpOnly: true,
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/login.html"))

	if r.Method == "GET" {
		tmpl.Execute(w, nil)
		return
	}

	r.ParseForm()
	username := strings.TrimSpace(r.FormValue("username"))
	password := strings.TrimSpace(r.FormValue("password"))

	var user models.User
	if err := storage.DB.Where("username = ?", username).First(&user).Error; err != nil {
		tmpl.Execute(w, map[string]string{"Error": "Неверный логин или пароль"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		tmpl.Execute(w, map[string]string{"Error": "Неверный логин или пароль"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	})
	tokenStr, _ := token.SignedString([]byte(jwtSecret))

	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    tokenStr,
		Path:     "/",
		MaxAge:   86400,
		HttpOnly: true,
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	username := getUsernameFromCookie(r)
	if username == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	user := storage.GetUserByUsername(username)
	if user == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	data := struct {
		Username string
		User     *models.User
		Success  string
		Error    string
	}{
		Username: username,
		User:     user,
		Success:  r.URL.Query().Get("success"),
		Error:    r.URL.Query().Get("error"),
	}

	tmpl, err := template.ParseFiles("templates/profile.html")
	if err != nil {
		http.Error(w, "Ошибка шаблона: "+err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}

func ProfileUpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.NotFound(w, r)
		return
	}

	username := getUsernameFromCookie(r)
	if username == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	user := storage.GetUserByUsername(username)
	if user == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	r.ParseForm()
	fullName := r.FormValue("fullName")
	email := r.FormValue("email")
	bio := r.FormValue("bio")
	location := r.FormValue("location")

	err := storage.UpdateUserProfile(user.ID, fullName, email, bio, location)
	if err != nil {
		http.Redirect(w, r, "/profile?error=Ошибка+обновления", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/profile?success=Профиль+обновлен", http.StatusSeeOther)
}

func ProfileChangePasswordHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.NotFound(w, r)
		return
	}

	username := getUsernameFromCookie(r)
	if username == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	user := storage.GetUserByUsername(username)
	if user == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	r.ParseForm()
	oldPassword := r.FormValue("oldPassword")
	newPassword := r.FormValue("newPassword")
	confirmPassword := r.FormValue("confirmPassword")

	if newPassword != confirmPassword {
		http.Redirect(w, r, "/profile?error=Пароли+не+совпадают", http.StatusSeeOther)
		return
	}

	if len(newPassword) < 6 {
		http.Redirect(w, r, "/profile?error=Пароль+должен+быть+минимум+6+символов", http.StatusSeeOther)
		return
	}

	_, err := storage.LoginUser(username, oldPassword)
	if err != nil {
		http.Redirect(w, r, "/profile?error=Неверный+текущий+пароль", http.StatusSeeOther)
		return
	}

	err = storage.UpdatePassword(user.ID, newPassword)
	if err != nil {
		http.Redirect(w, r, "/profile?error=Ошибка+смены+пароля", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/profile?success=Пароль+изменен", http.StatusSeeOther)
}
