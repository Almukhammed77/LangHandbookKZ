package main

import (
	"encoding/json"
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

// JavaScript Tutorials
var jsTutorials = []TutorialSection{
	{
		Title: "JavaScript Home",
		Slug:  "js-home",
		Content: `<h2>Добро пожаловать в JavaScript</h2>
		<p>JavaScript — это высокоуровневый язык программирования, который делает веб-страницы интерактивными. Создан Бренданом Эйхом в 1995 году всего за 10 дней!</p>
		
		<div style="background: linear-gradient(135deg, #667eea 0%, #764ba2 100%); color: white; padding: 25px; border-radius: 15px; margin: 20px 0;">
			<h3 style="color: white; border-bottom: 2px solid rgba(255,255,255,0.3); padding-bottom: 10px;">🌟 JavaScript сегодня</h3>
			<ul style="list-style-type: none; padding: 0;">
				<li style="margin: 10px 0;">✓ 98% всех веб-сайтов используют JavaScript</li>
				<li style="margin: 10px 0;">✓ Более 15 миллионов разработчиков по всему миру</li>
				<li style="margin: 10px 0;">✓ 1.5+ миллиона пакетов в npm</li>
				<li style="margin: 10px 0;">✓ Используется в Netflix, Uber, LinkedIn, eBay</li>
			</ul>
		</div>
		
		<h3>Что можно создавать с JavaScript?</h3>
		<table style="width:100%; border-collapse: collapse; margin: 20px 0;">
			<tr style="background: #f3f4f6;">
				<th style="padding: 12px; border: 1px solid #ddd; text-align: left;">Область</th>
				<th style="padding: 12px; border: 1px solid #ddd; text-align: left;">Фреймворки/Инструменты</th>
				<th style="padding: 12px; border: 1px solid #ddd; text-align: left;">Примеры</th>
			</tr>
			<tr>
				<td style="padding: 12px; border: 1px solid #ddd;">Frontend (веб)</td>
				<td style="padding: 12px; border: 1px solid #ddd;">React, Vue, Angular, Svelte</td>
				<td style="padding: 12px; border: 1px solid #ddd;">Facebook, Instagram, Gmail</td>
			</tr>
			<tr style="background: #f9f9f9;">
				<td style="padding: 12px; border: 1px solid #ddd;">Backend</td>
				<td style="padding: 12px; border: 1px solid #ddd;">Node.js, Deno, Bun</td>
				<td style="padding: 12px; border: 1px solid #ddd;">Netflix, PayPal, Uber</td>
			</tr>
			<tr>
				<td style="padding: 12px; border: 1px solid #ddd;">Мобильные приложения</td>
				<td style="padding: 12px; border: 1px solid #ddd;">React Native, Ionic, NativeScript</td>
				<td style="padding: 12px; border: 1px solid #ddd;">Instagram, Discord, Pinterest</td>
			</tr>
			<tr style="background: #f9f9f9;">
				<td style="padding: 12px; border: 1px solid #ddd;">Десктопные приложения</td>
				<td style="padding: 12px; border: 1px solid #ddd;">Electron, NW.js</td>
				<td style="padding: 12px; border: 1px solid #ddd;">VS Code, Slack, Discord, Figma</td>
			</tr>
		</table>`,
	},
	{
		Title: "JS Introduction",
		Slug:  "js-introduction",
		Content: `<h2>Введение в JavaScript</h2>
		
		<div style="background: #e8f5e9; padding: 20px; border-radius: 12px; margin: 20px 0;">
			<h3>📌 Ключевые особенности JavaScript</h3>
			<ul>
				<li><strong>Интерпретируемый язык</strong> — не требует компиляции, выполняется браузером</li>
				<li><strong>Динамическая типизация</strong> — переменные могут менять тип</li>
				<li><strong>Объектно-ориентированный</strong> — поддерживает ООП, но основан на прототипах</li>
				<li><strong>Функциональный</strong> — функции — объекты первого класса</li>
				<li><strong>Событийно-ориентированный</strong> — отлично подходит для интерактивных приложений</li>
			</ul>
		</div>
		
		<h3>История версий JavaScript</h3>
		<div style="display: grid; grid-template-columns: repeat(auto-fit, minmax(250px, 1fr)); gap: 15px; margin: 20px 0;">
			<div style="background: #f3f4f6; padding: 15px; border-radius: 10px;">
				<h4 style="margin-top: 0;">ES5 (2009)</h4>
				<p>Поддержка JSON, строгий режим ('use strict'), новые методы массивов (forEach, map, filter)</p>
			</div>
			<div style="background: #e0f2fe; padding: 15px; border-radius: 10px;">
				<h4 style="margin-top: 0;">ES6/ES2015</h4>
				<p>Революционное обновление: let/const, стрелочные функции, классы, промисы, модули</p>
			</div>
			<div style="background: #f3e5f5; padding: 15px; border-radius: 10px;">
				<h4 style="margin-top: 0;">ES2016-ES2023</h4>
				<p>Async/await, операторы расширения, nullish coalescing, приватные поля</p>
			</div>
		</div>`,
	},
	{
		Title: "JS Syntax",
		Slug:  "js-syntax",
		Content: `<h2>Синтаксис JavaScript</h2>
		
		<h3>Основные правила синтаксиса</h3>
		<ul>
			<li>JavaScript чувствителен к регистру</li>
			<li>Инструкции обычно заканчиваются точкой с запятой (;)</li>
			<li>Комментарии: // однострочные, /* многострочные */</li>
		</ul>
		
		<h3>Переменные</h3>
		<pre><code>// var (устаревший)
var x = 5;

// let (изменяемая)
let age = 25;
age = 26; // можно изменять

// const (неизменяемая)
const PI = 3.14159;</code></pre>
		
		<h3>Типы данных</h3>
		<table style="width:100%; border-collapse: collapse; margin: 15px 0;">
			<tr style="background: #f3f4f6;">
				<th style="padding: 10px; border: 1px solid #ddd;">Тип</th>
				<th style="padding: 10px; border: 1px solid #ddd;">Пример</th>
			</tr>
			<tr><td style="padding: 10px; border: 1px solid #ddd;">string</td><td style="padding: 10px; border: 1px solid #ddd;">"hello", 'world'</td></tr>
			<tr><td style="padding: 10px; border: 1px solid #ddd;">number</td><td style="padding: 10px; border: 1px solid #ddd;">42, 3.14</td></tr>
			<tr><td style="padding: 10px; border: 1px solid #ddd;">boolean</td><td style="padding: 10px; border: 1px solid #ddd;">true, false</td></tr>
			<tr><td style="padding: 10px; border: 1px solid #ddd;">object</td><td style="padding: 10px; border: 1px solid #ddd;">{name: "Азамат"}</td></tr>
			<tr><td style="padding: 10px; border: 1px solid #ddd;">array</td><td style="padding: 10px; border: 1px solid #ddd;">[1, 2, 3]</td></tr>
			<tr><td style="padding: 10px; border: 1px solid #ddd;">undefined</td><td style="padding: 10px; border: 1px solid #ddd;">let x;</td></tr>
			<tr><td style="padding: 10px; border: 1px solid #ddd;">null</td><td style="padding: 10px; border: 1px solid #ddd;">let x = null;</td></tr>
		</table>`,
	},
	{
		Title: "JS Functions",
		Slug:  "js-functions",
		Content: `<h2>Функции в JavaScript</h2>
		
		<h3>Function Declaration</h3>
		<pre><code>function greet(name) {
    return "Привет, " + name + "!";
}</code></pre>
		
		<h3>Function Expression</h3>
		<pre><code>const greet = function(name) {
    return "Привет, " + name + "!";
};</code></pre>
		
		<h3>Arrow Functions</h3>
		<pre><code>const greet = (name) => "Привет, " + name + "!";
const square = x => x * x;</code></pre>
		
		<h3>Параметры по умолчанию</h3>
		<pre><code>function greet(name = "гость") {
    return "Привет, " + name;
}</code></pre>`,
	},
	{
		Title: "JS Arrays",
		Slug:  "js-arrays",
		Content: `<h2>Массивы в JavaScript</h2>
		
		<h3>Создание массивов</h3>
		<pre><code>let fruits = ["Яблоко", "Банан", "Апельсин"];
let numbers = new Array(1, 2, 3, 4, 5);</code></pre>
		
		<h3>Основные методы</h3>
		<pre><code>fruits.push("Груша");      // Добавить в конец
fruits.pop();               // Удалить с конца
fruits.unshift("Вишня");    // Добавить в начало
fruits.shift();             // Удалить с начала
fruits.indexOf("Банан");    // Индекс элемента
fruits.includes("Яблоко");  // Проверка наличия</code></pre>
		
		<h3>Методы для перебора</h3>
		<pre><code>// forEach
fruits.forEach(fruit => console.log(fruit));

// map
let lengths = fruits.map(fruit => fruit.length);

// filter
let longNames = fruits.filter(fruit => fruit.length > 5);

// reduce
let total = [1,2,3,4,5].reduce((sum, num) => sum + num, 0);</code></pre>`,
	},
	{
		Title: "JS Objects",
		Slug:  "js-objects",
		Content: `<h2>Объекты в JavaScript</h2>
		
		<h3>Создание объектов</h3>
		<pre><code>const user = {
    name: "Азамат",
    age: 25,
    city: "Алматы",
    greet() {
        return "Привет, я " + this.name;
    }
};</code></pre>
		
		<h3>Доступ к свойствам</h3>
		<pre><code>console.log(user.name);      // Точечная нотация
console.log(user["age"]);    // Скобочная нотация</code></pre>
		
		<h3>Методы объектов</h3>
		<pre><code>Object.keys(user)    // ["name", "age", "city"]
Object.values(user)  // ["Азамат", 25, "Алматы"]
Object.entries(user) // [["name","Азамат"], ...]</code></pre>
		
		<h3>Деструктуризация</h3>
		<pre><code>const {name, age} = user;
console.log(name); // "Азамат"</code></pre>`,
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
element.innerHTML = "<strong>Жирный текст</strong>";
element.style.color = "red";</code></pre>
		
		<h3>Работа с классами</h3>
		<pre><code>element.classList.add("active");
element.classList.remove("hidden");
element.classList.toggle("visible");</code></pre>
		
		<h3>События</h3>
		<pre><code>button.addEventListener("click", function() {
    console.log("Клик!");
});</code></pre>`,
	},
	{
		Title: "JS Async",
		Slug:  "js-async",
		Content: `<h2>Асинхронность в JavaScript</h2>
		
		<h3>Callbacks</h3>
		<pre><code>setTimeout(() => {
    console.log("Прошло 2 секунды");
}, 2000);</code></pre>
		
		<h3>Promises</h3>
		<pre><code>fetch("https://api.example.com")
    .then(res => res.json())
    .then(data => console.log(data))
    .catch(err => console.error(err));</code></pre>
		
		<h3>Async/Await</h3>
		<pre><code>async function getData() {
    try {
        const res = await fetch("https://api.example.com");
        const data = await res.json();
        console.log(data);
    } catch (err) {
        console.error(err);
    }
}</code></pre>`,
	},
}

// Go Tutorials
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
			<li><strong>Читаемость:</strong> единый стиль кода</li>
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
		<pre><code># macOS
brew install go

# Linux
sudo apt install golang-go

# Windows
# Скачайте с golang.org/dl</code></pre>
		
		<h3>Первая программа</h3>
		<pre><code>package main

import "fmt"

func main() {
    fmt.Println("Сәлем, Go!")
}</code></pre>
		
		<h3>Запуск программы</h3>
		<pre><code>go run main.go
go build main.go  # компиляция в бинарник</code></pre>`,
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
    city := "Астана"  // короткое объявление
    
    fmt.Printf("Привет, %s! Возраст: %d\n", name, age)
}</code></pre>
		
		<h3>Типы данных</h3>
		<pre><code>var i int = 42
var f float64 = 3.14
var s string = "Текст"
var b bool = true
var arr [3]int = [3]int{1, 2, 3}</code></pre>`,
	},
	{
		Title: "Go Variables",
		Slug:  "go-variables",
		Content: `<h2>Переменные в Go</h2>
		
		<h3>Объявление переменных</h3>
		<pre><code>var name string = "Азамат"
var city = "Алматы"  // вывод типа
age := 25            // короткое объявление</code></pre>
		
		<h3>Множественное объявление</h3>
		<pre><code>var x, y int = 1, 2
a, b := "hello", true</code></pre>
		
		<h3>Константы</h3>
		<pre><code>const Pi = 3.14159
const StatusOK = 200</code></pre>
		
		<h3>Нулевые значения</h3>
		<pre><code>var i int     // 0
var f float64 // 0
var s string  // ""
var b bool    // false
var p *int    // nil</code></pre>`,
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
}</code></pre>
		
		<h3>Именованные возвращаемые значения</h3>
		<pre><code>func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}</code></pre>`,
	},
	{
		Title: "Go Arrays",
		Slug:  "go-arrays",
		Content: `<h2>Массивы и Срезы</h2>
		
		<h3>Массивы</h3>
		<pre><code>var arr [5]int
arr[0] = 1
arr[1] = 2

primes := [5]int{2, 3, 5, 7, 11}</code></pre>
		
		<h3>Срезы (Slices)</h3>
		<pre><code>slice := []int{1, 2, 3}
slice = append(slice, 4, 5)

// Срез из массива
arr := [5]int{1, 2, 3, 4, 5}
s := arr[1:4] // [2, 3, 4]</code></pre>
		
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
		<pre><code>// Литерал
user := map[string]string{
    "name": "Азамат",
    "city": "Алматы",
}

// С помощью make
ages := make(map[string]int)
ages["Азамат"] = 25</code></pre>
		
		<h3>Операции с map</h3>
		<pre><code>// Получение значения
name := user["name"]

// Проверка существования
value, exists := user["phone"]

// Удаление
delete(user, "age")

// Длина
len(user)</code></pre>`,
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
		<pre><code>user1 := User{ID: 1, Name: "Айгерим", Age: 25}
user2 := User{Name: "Азамат"} // остальные поля - нулевые</code></pre>
		
		<h3>Методы структур</h3>
		<pre><code>func (u User) Greet() string {
    return "Привет, я " + u.Name
}

func (u *User) Birthday() {
    u.Age++
}</code></pre>
		
		<h3>Теги структур</h3>
		<pre><code>type Config struct {
    Host string ` + "`json:\"host\"`" + `
    Port int    ` + "`json:\"port\" default:\"8080\"`" + `
}</code></pre>`,
	},
}

// TypeScript Tutorials
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
				<li>Самодокументируемый код</li>
			</ul>
		</div>
		
		<h3>Кто использует TypeScript?</h3>
		<ul>
			<li>Google (Angular)</li>
			<li>Microsoft (VS Code)</li>
			<li>Airbnb</li>
			<li>Slack</li>
		</ul>`,
	},
	{
		Title: "TS Introduction",
		Slug:  "ts-introduction",
		Content: `<h2>Введение в TypeScript</h2>
		
		<h3>Что такое TypeScript?</h3>
		<p>TypeScript добавляет в JavaScript статическую типизацию, классы, интерфейсы, дженерики.</p>
		
		<h3>TypeScript vs JavaScript</h3>
		<table style="width:100%; border-collapse: collapse;">
			<tr style="background: #f3f4f6;">
				<th style="padding: 10px; border: 1px solid #ddd;">TypeScript</th>
				<th style="padding: 10px; border: 1px solid #ddd;">JavaScript</th>
			</tr>
			<tr>
				<td style="padding: 10px; border: 1px solid #ddd;">Статическая типизация</td>
				<td style="padding: 10px; border: 1px solid #ddd;">Динамическая типизация</td>
			</tr>
			<tr>
				<td style="padding: 10px; border: 1px solid #ddd;">Ошибки на этапе компиляции</td>
				<td style="padding: 10px; border: 1px solid #ddd;">Ошибки в рантайме</td>
			</tr>
			<tr>
				<td style="padding: 10px; border: 1px solid #ddd;">Поддержка интерфейсов</td>
				<td style="padding: 10px; border: 1px solid #ddd;">Нет интерфейсов</td>
			</tr>
		</table>`,
	},
	{
		Title: "TS Basic Types",
		Slug:  "ts-basic-types",
		Content: `<h2>Базовые типы TypeScript</h2>
		
		<h3>Примитивные типы</h3>
		<pre><code>let name: string = "Алматы";
let age: number = 28;
let isActive: boolean = true;
let numbers: number[] = [1, 2, 3];
let tuple: [string, number] = ["hello", 42];</code></pre>
		
		<h3>any и unknown</h3>
		<pre><code>let dynamic: any = 4;
dynamic = "string"; // OK

let safe: unknown = "hello";
if (typeof safe === "string") {
    console.log(safe.toUpperCase());
}</code></pre>
		
		<h3>void и never</h3>
		<pre><code>function log(message: string): void {
    console.log(message);
}

function error(message: string): never {
    throw new Error(message);
}</code></pre>`,
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
}</code></pre>
		
		<h3>Параметры по умолчанию</h3>
		<pre><code>function greet(name: string = "гость"): string {
    return "Привет, " + name;
}</code></pre>
		
		<h3>Перегрузка функций</h3>
		<pre><code>function reverse(str: string): string;
function reverse(arr: number[]): number[];
function reverse(value: string | number[]): any {
    // реализация
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
    email?: string;  // опциональное поле
    readonly createdAt: Date;
}

const user: User = {
    id: 1,
    name: "Айжан",
    createdAt: new Date()
};</code></pre>
		
		<h3>Расширение интерфейсов</h3>
		<pre><code>interface Person {
    name: string;
    age: number;
}

interface Employee extends Person {
    employeeId: number;
    department: string;
}</code></pre>`,
	},
	{
		Title: "TS Classes",
		Slug:  "ts-classes",
		Content: `<h2>Классы в TypeScript</h2>
		
		<h3>Модификаторы доступа</h3>
		<pre><code>class Person {
    public name: string;      // доступно всем
    private age: number;       // только внутри класса
    protected email: string;   // внутри класса и наследников
    readonly id: number;       // только для чтения
    
    constructor(name: string, age: number, email: string, id: number) {
        this.name = name;
        this.age = age;
        this.email = email;
        this.id = id;
    }
}</code></pre>
		
		<h3>Сокращенная инициализация</h3>
		<pre><code>class User {
    constructor(
        public name: string,
        private age: number,
        readonly id: number
    ) {}
}</code></pre>`,
	},
	{
		Title: "TS Generics",
		Slug:  "ts-generics",
		Content: `<h2>Дженерики в TypeScript</h2>
		
		<h3>Generic функции</h3>
		<pre><code>function identity<T>(arg: T): T {
    return arg;
}

let output = identity<string>("hello");</code></pre>
		
		<h3>Generic интерфейсы</h3>
		<pre><code>interface Box<T> {
    value: T;
    getValue(): T;
}

class StringBox implements Box<string> {
    constructor(private _value: string) {}
    
    get value(): string { return this._value; }
    getValue(): string { return this._value; }
}</code></pre>
		
		<h3>Generic ограничения</h3>
		<pre><code>interface Lengthwise {
    length: number;
}

function logLength<T extends Lengthwise>(arg: T): T {
    console.log(arg.length);
    return arg;
}</code></pre>`,
	},
	{
		Title: "TS Modules",
		Slug:  "ts-modules",
		Content: `<h2>Модули в TypeScript</h2>
		
		<h3>Экспорт</h3>
		<pre><code>// math.ts
export function add(a: number, b: number): number {
    return a + b;
}

export const PI = 3.14159;

export default class Calculator {
    multiply(a: number, b: number): number {
        return a * b;
    }
}</code></pre>
		
		<h3>Импорт</h3>
		<pre><code>// app.ts
import { add, PI } from "./math.js";
import Calculator from "./math.js";
import * as math from "./math.js";</code></pre>`,
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

// Partial — все поля опциональные
type PartialUser = Partial<User>;

// Readonly — все поля только для чтения
type ReadonlyUser = Readonly<User>;

// Pick — выбирает указанные поля
type UserPreview = Pick<User, "id" | "name">;

// Omit — исключает указанные поля
type UserWithoutEmail = Omit<User, "email">;

// Record — словарь
type UserRoles = Record<number, string>;</code></pre>`,
	},
}

// Python Tutorials
var pythonTutorials = []TutorialSection{
	{
		Title: "Python Home",
		Slug:  "python-home",
		Content: `<h2>Добро пожаловать в Python!</h2>
		<p>Python — самый популярный язык для начинающих. Создан Гвидо ван Россумом в 1991 году.</p>
		
		<div style="background: #e6f7e6; padding: 20px; border-radius: 12px; margin: 20px 0;">
			<h3>Почему Python?</h3>
			<ul>
				<li>Простой и понятный синтаксис</li>
				<li>Огромное количество библиотек</li>
				<li>Используется в AI, Data Science, Web</li>
				<li>Кроссплатформенность</li>
			</ul>
		</div>
		
		<h3>Где используется Python?</h3>
		<ul>
			<li>Веб-разработка (Django, Flask)</li>
			<li>Машинное обучение (TensorFlow, PyTorch)</li>
			<li>Анализ данных (Pandas, NumPy)</li>
			<li>Автоматизация и скриптинг</li>
			<li>Научные вычисления</li>
		</ul>`,
	},
	{
		Title: "Python Introduction",
		Slug:  "python-introduction",
		Content: `<h2>Введение в Python</h2>
		
		<h3>Установка Python</h3>
		<pre><code># macOS
brew install python

# Linux
sudo apt install python3 python3-pip

# Windows
# Скачайте с python.org</code></pre>
		
		<h3>Первая программа</h3>
		<pre><code>print("Сәлем, Python!")
print("Привет из Казахстана!")</code></pre>
		
		<h3>Особенности Python</h3>
		<ul>
			<li>Динамическая типизация</li>
			<li>Интерпретируемый язык</li>
			<li>Автоматическое управление памятью</li>
			<li>Отступы вместо скобок</li>
		</ul>`,
	},
	{
		Title: "Python Syntax",
		Slug:  "python-syntax",
		Content: `<h2>Синтаксис Python</h2>
		
		<h3>Основные правила</h3>
		<ul>
			<li>Отступы важны! Обычно 4 пробела</li>
			<li>Комментарии начинаются с #</li>
			<li>Имена переменных чувствительны к регистру</li>
		</ul>
		
		<h3>Переменные</h3>
		<pre><code>name = "Азамат"
age = 25
height = 1.75
is_student = True</code></pre>
		
		<h3>Ввод и вывод</h3>
		<pre><code>name = input("Введите имя: ")
print(f"Привет, {name}!")</code></pre>`,
	},
	{
		Title: "Python Variables",
		Slug:  "python-variables",
		Content: `<h2>Переменные в Python</h2>
		
		<h3>Объявление переменных</h3>
		<pre><code># Простое присваивание
name = "Азамат"
age = 25

# Множественное присваивание
x, y, z = 1, 2, 3
a = b = c = 0</code></pre>
		
		<h3>Типы данных</h3>
		<table style="width:100%; border-collapse: collapse;">
			<tr style="background: #f3f4f6;">
				<th style="padding: 10px; border: 1px solid #ddd;">Тип</th>
				<th style="padding: 10px; border: 1px solid #ddd;">Пример</th>
			</tr>
			<tr><td style="padding: 10px; border: 1px solid #ddd;">int</td><td style="padding: 10px; border: 1px solid #ddd;">42</td></tr>
			<tr><td style="padding: 10px; border: 1px solid #ddd;">float</td><td style="padding: 10px; border: 1px solid #ddd;">3.14</td></tr>
			<tr><td style="padding: 10px; border: 1px solid #ddd;">str</td><td style="padding: 10px; border: 1px solid #ddd;">"hello"</td></tr>
			<tr><td style="padding: 10px; border: 1px solid #ddd;">bool</td><td style="padding: 10px; border: 1px solid #ddd;">True, False</td></tr>
			<tr><td style="padding: 10px; border: 1px solid #ddd;">list</td><td style="padding: 10px; border: 1px solid #ddd;">[1, 2, 3]</td></tr>
			<tr><td style="padding: 10px; border: 1px solid #ddd;">tuple</td><td style="padding: 10px; border: 1px solid #ddd;">(1, 2, 3)</td></tr>
			<tr><td style="padding: 10px; border: 1px solid #ddd;">dict</td><td style="padding: 10px; border: 1px solid #ddd;">{"key": "value"}</td></tr>
			<tr><td style="padding: 10px; border: 1px solid #ddd;">set</td><td style="padding: 10px; border: 1px solid #ddd;">{1, 2, 3}</td></tr>
		</table>`,
	},
	{
		Title: "Python Data Types",
		Slug:  "python-datatypes",
		Content: `<h2>Типы данных в Python</h2>
		
		<h3>Строки (str)</h3>
		<pre><code>s1 = 'single quotes'
s2 = "double quotes"
s3 = """multi line"""

# Методы строк
text = "  Hello, World!  "
text.lower()           # "  hello, world!  "
text.upper()           # "  HELLO, WORLD!  "
text.strip()           # "Hello, World!"
text.split(",")        # ["  Hello", " World!  "]
",".join(["a", "b"])   # "a,b"</code></pre>
		
		<h3>Списки (list)</h3>
		<pre><code>fruits = ["🍎", "🍌", "🍊"]
fruits.append("🍇")
fruits.insert(1, "🍓")
fruits.remove("🍌")
popped = fruits.pop()</code></pre>
		
		<h3>Словари (dict)</h3>
		<pre><code>user = {
    "name": "Азамат",
    "age": 25,
    "city": "Алматы"
}

user["email"] = "azamat@example.com"
del user["age"]</code></pre>`,
	},
	{
		Title: "Python If...Else",
		Slug:  "python-ifelse",
		Content: `<h2>Условные операторы</h2>
		
		<h3>if, elif, else</h3>
		<pre><code>age = 18

if age >= 18:
    print("Взрослый")
elif age >= 13:
    print("Подросток")
else:
    print("Ребенок")</code></pre>
		
		<h3>Тернарный оператор</h3>
		<pre><code>status = "Взрослый" if age >= 18 else "Несовершеннолетний"</code></pre>
		
		<h3>Логические операторы</h3>
		<pre><code>if age >= 18 and has_license:
    print("Можно водить")

if is_weekend or is_holiday:
    print("Можно отдыхать")</code></pre>`,
	},
	{
		Title: "Python Loops",
		Slug:  "python-loops",
		Content: `<h2>Циклы в Python</h2>
		
		<h3>Цикл for</h3>
		<pre><code># По списку
fruits = ["🍎", "🍌", "🍊"]
for fruit in fruits:
    print(fruit)

# По диапазону
for i in range(5):
    print(i)

# С индексом
for i, fruit in enumerate(fruits):
    print(f"{i}: {fruit}")</code></pre>
		
		<h3>Цикл while</h3>
		<pre><code>count = 0
while count < 5:
    print(count)
    count += 1</code></pre>
		
		<h3>Управление циклами</h3>
		<pre><code># break - выход из цикла
for i in range(10):
    if i == 5:
        break

# continue - пропуск итерации
for i in range(5):
    if i == 2:
        continue</code></pre>`,
	},
}

// Java Tutorials
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
				<li>92% enterprise-приложений используют Java</li>
			</ul>
		</div>
		
		<h3>Основные области применения</h3>
		<ul>
			<li>Android-приложения</li>
			<li>Веб-приложения (Spring)</li>
			<li>Enterprise-системы</li>
			<li>Научные приложения</li>
		</ul>`,
	},
	{
		Title: "Java Introduction",
		Slug:  "java-introduction",
		Content: `<h2>Введение в Java</h2>
		
		<h3>Установка Java</h3>
		<pre><code># Проверка установки
java -version
javac -version

# Компиляция и запуск
javac HelloWorld.java
java HelloWorld</code></pre>
		
		<h3>Первая программа</h3>
		<pre><code>public class HelloWorld {
    public static void main(String[] args) {
        System.out.println("Сәлем, Java!");
    }
}</code></pre>
		
		<h3>Особенности Java</h3>
		<ul>
			<li>Строгая статическая типизация</li>
			<li>Объектно-ориентированный</li>
			<li>Автоматическое управление памятью</li>
			<li>Платформонезависимость (JVM)</li>
		</ul>`,
	},
	{
		Title: "Java Syntax",
		Slug:  "java-syntax",
		Content: `<h2>Синтаксис Java</h2>
		
		<h3>Структура класса</h3>
		<pre><code>public class Main {
    public static void main(String[] args) {
        System.out.println("Hello, World!");
    }
}</code></pre>
		
		<h3>Комментарии</h3>
		<pre><code>// Однострочный комментарий

/*
   Многострочный
   комментарий
*/

/**
 * JavaDoc комментарий
 */</code></pre>
		
		<h3>Точка с запятой</h3>
		<p>Каждая инструкция в Java должна заканчиваться точкой с запятой (;).</p>`,
	},
	{
		Title: "Java Variables",
		Slug:  "java-variables",
		Content: `<h2>Переменные в Java</h2>
		
		<h3>Объявление переменных</h3>
		<pre><code>int age = 25;
double price = 99.99;
String name = "Азамат";
boolean isActive = true;
char grade = 'A';</code></pre>
		
		<h3>Типы данных</h3>
		<table style="width:100%; border-collapse: collapse;">
			<tr style="background: #f3f4f6;">
				<th style="padding: 10px; border: 1px solid #ddd;">Тип</th>
				<th style="padding: 10px; border: 1px solid #ddd;">Размер</th>
				<th style="padding: 10px; border: 1px solid #ddd;">Диапазон</th>
			</tr>
			<tr><td style="padding: 10px; border: 1px solid #ddd;">byte</td><td style="padding: 10px; border: 1px solid #ddd;">1 байт</td><td style="padding: 10px; border: 1px solid #ddd;">-128 до 127</td></tr>
			<tr><td style="padding: 10px; border: 1px solid #ddd;">short</td><td style="padding: 10px; border: 1px solid #ddd;">2 байта</td><td style="padding: 10px; border: 1px solid #ddd;">-32,768 до 32,767</td></tr>
			<tr><td style="padding: 10px; border: 1px solid #ddd;">int</td><td style="padding: 10px; border: 1px solid #ddd;">4 байта</td><td style="padding: 10px; border: 1px solid #ddd;">-2^31 до 2^31-1</td></tr>
			<tr><td style="padding: 10px; border: 1px solid #ddd;">long</td><td style="padding: 10px; border: 1px solid #ddd;">8 байт</td><td style="padding: 10px; border: 1px solid #ddd;">-2^63 до 2^63-1</td></tr>
			<tr><td style="padding: 10px; border: 1px solid #ddd;">float</td><td style="padding: 10px; border: 1px solid #ddd;">4 байта</td><td style="padding: 10px; border: 1px solid #ddd;">~±3.4E+38</td></tr>
			<tr><td style="padding: 10px; border: 1px solid #ddd;">double</td><td style="padding: 10px; border: 1px solid #ddd;">8 байт</td><td style="padding: 10px; border: 1px solid #ddd;">~±1.8E+308</td></tr>
		</table>`,
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
int div = a / b;      // 3 (целочисленное деление)
int mod = a % b;      // 1 (остаток)</code></pre>
		
		<h3>Операторы сравнения</h3>
		<pre><code>==, !=, <, >, <=, >=</code></pre>
		
		<h3>Логические операторы</h3>
		<pre><code>&& (AND), || (OR), ! (NOT)</code></pre>
		
		<h3>Инкремент/Декремент</h3>
		<pre><code>a++;  // постфиксный
++a;  // префиксный
b--;
--b;</code></pre>`,
	},
	{
		Title: "Java Control Flow",
		Slug:  "java-control-flow",
		Content: `<h2>Управляющие конструкции</h2>
		
		<h3>if-else</h3>
		<pre><code>if (age >= 18) {
    System.out.println("Взрослый");
} else if (age >= 13) {
    System.out.println("Подросток");
} else {
    System.out.println("Ребенок");
}</code></pre>
		
		<h3>switch</h3>
		<pre><code>switch (day) {
    case 1:
        System.out.println("Понедельник");
        break;
    case 2:
        System.out.println("Вторник");
        break;
    default:
        System.out.println("Неизвестный день");
}</code></pre>`,
	},
	{
		Title: "Java Loops",
		Slug:  "java-loops",
		Content: `<h2>Циклы в Java</h2>
		
		<h3>for loop</h3>
		<pre><code>for (int i = 0; i < 5; i++) {
    System.out.println(i);
}

// enhanced for
int[] numbers = {1, 2, 3, 4, 5};
for (int num : numbers) {
    System.out.println(num);
}</code></pre>
		
		<h3>while loop</h3>
		<pre><code>int i = 0;
while (i < 5) {
    System.out.println(i);
    i++;
}</code></pre>
		
		<h3>do-while</h3>
		<pre><code>int i = 0;
do {
    System.out.println(i);
    i++;
} while (i < 5);</code></pre>`,
	},
	{
		Title: "Java Arrays",
		Slug:  "java-arrays",
		Content: `<h2>Массивы в Java</h2>
		
		<h3>Объявление массивов</h3>
		<pre><code>// Объявление и создание
int[] numbers = new int[5];
numbers[0] = 1;
numbers[1] = 2;

// Инициализация значениями
int[] arr = {1, 2, 3, 4, 5};

// Многомерные массивы
int[][] matrix = {
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9}
};</code></pre>
		
		<h3>Длина массива</h3>
		<pre><code>int length = numbers.length;</code></pre>
		
		<h3>Копирование массивов</h3>
		<pre><code>int[] copy = Arrays.copyOf(arr, arr.length);
int[] range = Arrays.copyOfRange(arr, 1, 4);</code></pre>`,
	},
}

// C# Tutorials
var csharpTutorials = []TutorialSection{
	{
		Title: "C# Home",
		Slug:  "csharp-home",
		Content: `<h2>Добро пожаловать в C#!</h2>
		<p>C# — современный объектно-ориентированный язык от Microsoft, созданный в 2000 году Андерсом Хейлсбергом.</p>
		
		<div style="background: #e6e6fa; padding: 20px; border-radius: 12px; margin: 20px 0;">
			<h3>Где используется C#?</h3>
			<ul>
				<li>Разработка игр (Unity)</li>
				<li>Веб-приложения (ASP.NET)</li>
				<li>Десктопные приложения (WPF, WinForms)</li>
				<li>Мобильные приложения (Xamarin)</li>
				<li>Облачные приложения (Azure)</li>
			</ul>
		</div>
		
		<h3>Преимущества C#</h3>
		<ul>
			<li>Строгая типизация</li>
			<li>Современный синтаксис</li>
			<li>Отличная интеграция с Windows</li>
			<li>Кроссплатформенность (.NET Core)</li>
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
}</code></pre>
		
		<h3>Структура программы</h3>
		<ul>
			<li>Пространства имен (namespace)</li>
			<li>Классы (class)</li>
			<li>Методы (method)</li>
			<li>Операторы (statements)</li>
		</ul>`,
	},
	{
		Title: "C# Syntax",
		Slug:  "csharp-syntax",
		Content: `<h2>Синтаксис C#</h2>
		
		<h3>Переменные</h3>
		<pre><code>int age = 25;
string name = "Азамат";
bool isActive = true;
double price = 99.99;
decimal salary = 5000.50m;</code></pre>
		
		<h3>Вывод в консоль</h3>
		<pre><code>Console.WriteLine("Текст с переносом");
Console.Write("Текст без переноса");
Console.WriteLine($"Имя: {name}, Возраст: {age}");</code></pre>
		
		<h3>Ввод с консоли</h3>
		<pre><code>string input = Console.ReadLine();
int number = Convert.ToInt32(input);</code></pre>`,
	},
	{
		Title: "C# Operators",
		Slug:  "csharp-operators",
		Content: `<h2>Операторы в C#</h2>
		
		<h3>Арифметические</h3>
		<pre><code>int a = 10, b = 3;
int sum = a + b;
int diff = a - b;
int mult = a * b;
int div = a / b;
int mod = a % b;</code></pre>
		
		<h3>Операторы сравнения</h3>
		<pre><code>==, !=, <, >, <=, >=</code></pre>
		
		<h3>Логические</h3>
		<pre><code>&&, ||, !</code></pre>
		
		<h3>Условный оператор</h3>
		<pre><code>string result = (age >= 18) ? "Взрослый" : "Несовершеннолетний";</code></pre>`,
	},
	{
		Title: "C# Control Flow",
		Slug:  "csharp-control-flow",
		Content: `<h2>Управляющие конструкции</h2>
		
		<h3>if-else</h3>
		<pre><code>if (age >= 18) {
    Console.WriteLine("Взрослый");
} else if (age >= 13) {
    Console.WriteLine("Подросток");
} else {
    Console.WriteLine("Ребенок");
}</code></pre>
		
		<h3>switch</h3>
		<pre><code>switch (day) {
    case 1:
        Console.WriteLine("Понедельник");
        break;
    case 2:
        Console.WriteLine("Вторник");
        break;
    default:
        Console.WriteLine("Неизвестный день");
        break;
}</code></pre>`,
	},
	{
		Title: "C# Loops",
		Slug:  "csharp-loops",
		Content: `<h2>Циклы в C#</h2>
		
		<h3>for</h3>
		<pre><code>for (int i = 0; i < 5; i++) {
    Console.WriteLine(i);
}</code></pre>
		
		<h3>foreach</h3>
		<pre><code>int[] numbers = {1, 2, 3, 4, 5};
foreach (int num in numbers) {
    Console.WriteLine(num);
}</code></pre>
		
		<h3>while</h3>
		<pre><code>int i = 0;
while (i < 5) {
    Console.WriteLine(i);
    i++;
}</code></pre>`,
	},
	{
		Title: "C# Arrays",
		Slug:  "csharp-arrays",
		Content: `<h2>Массивы в C#</h2>
		
		<h3>Объявление массивов</h3>
		<pre><code>// Одномерный массив
int[] numbers = new int[5];
int[] arr = {1, 2, 3, 4, 5};

// Многомерный массив
int[,] matrix = new int[3, 3];

// Зубчатый массив (jagged array)
int[][] jagged = new int[3][];
jagged[0] = new int[] {1, 2};</code></pre>
		
		<h3>Методы массивов</h3>
		<pre><code>Array.Sort(arr);
Array.Reverse(arr);
int index = Array.IndexOf(arr, 3);</code></pre>`,
	},
}

// Rust Tutorials
var rustTutorials = []TutorialSection{
	{
		Title: "Rust Home",
		Slug:  "rust-home",
		Content: `<h2>Добро пожаловать в Rust!</h2>
		<p>Rust — системный язык программирования от Mozilla, созданный в 2010 году. 7 лет подряд признается самым любимым языком по опросу Stack Overflow.</p>
		
		<div style="background: #ffdab9; padding: 20px; border-radius: 12px; margin: 20px 0;">
			<h3>Особенности Rust</h3>
			<ul>
				<li>Безопасность памяти без сборщика мусора</li>
				<li>Высокая производительность (как C/C++)</li>
				<li>Современный синтаксис</li>
				<li>Отличная документация</li>
				<li>Активное сообщество</li>
			</ul>
		</div>
		
		<h3>Где используется Rust?</h3>
		<ul>
			<li>Системное программирование</li>
			<li>Веб-ассемблер (WASM)</li>
			<li>Блокчейн проекты</li>
			<li>Игровые движки</li>
			<li>Операционные системы</li>
		</ul>`,
	},
	{
		Title: "Rust Introduction",
		Slug:  "rust-introduction",
		Content: `<h2>Введение в Rust</h2>
		
		<h3>Установка Rust</h3>
		<pre><code>curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh</code></pre>
		
		<h3>Первая программа</h3>
		<pre><code>fn main() {
    println!("Сәлем, Rust!");
}</code></pre>
		
		<h3>Компиляция и запуск</h3>
		<pre><code>rustc main.rs
./main

# Или через cargo
cargo new my_project
cargo run</code></pre>`,
	},
	{
		Title: "Rust Variables",
		Slug:  "rust-variables",
		Content: `<h2>Переменные в Rust</h2>
		
		<h3>Неизменяемые переменные</h3>
		<pre><code>let x = 5;        // неизменяемая
// x = 6;         // ошибка!</code></pre>
		
		<h3>Изменяемые переменные</h3>
		<pre><code>let mut y = 5;    // изменяемая
y = 6;            // OK</code></pre>
		
		<h3>Константы</h3>
		<pre><code>const MAX_POINTS: u32 = 100_000;</code></pre>
		
		<h3>Затенение (shadowing)</h3>
		<pre><code>let x = 5;
let x = x + 1;    // новая переменная</code></pre>`,
	},
	{
		Title: "Rust Functions",
		Slug:  "rust-functions",
		Content: `<h2>Функции в Rust</h2>
		
		<h3>Объявление функций</h3>
		<pre><code>fn add(x: i32, y: i32) -> i32 {
    x + y  // без точки с запятой - возвращаемое значение
}

fn greet(name: &str) {
    println!("Привет, {}!", name);
}</code></pre>
		
		<h3>Выражения и операторы</h3>
		<pre><code>let y = {
    let x = 3;
    x + 1  // выражение, возвращает значение
};</code></pre>`,
	},
	{
		Title: "Rust Ownership",
		Slug:  "rust-ownership",
		Content: `<h2>Владение (Ownership)</h2>
		
		<h3>Правила владения</h3>
		<ul>
			<li>У каждого значения есть владелец</li>
			<li>Может быть только один владелец</li>
			<li>Когда владелец выходит из области видимости, значение удаляется</li>
		</ul>
		
		<pre><code>let s1 = String::from("hello");
let s2 = s1;  // владение перемещается в s2
// println!("{}", s1); // ошибка!

let s3 = s2.clone();  // глубокое копирование
println!("{}", s2);    // OK</code></pre>
		
		<h3>Заимствование</h3>
		<pre><code>fn calculate_length(s: &String) -> usize {  // & - заимствование
    s.len()
}</code></pre>`,
	},
	{
		Title: "Rust Structs",
		Slug:  "rust-structs",
		Content: `<h2>Структуры в Rust</h2>
		
		<h3>Определение структуры</h3>
		<pre><code>struct User {
    username: String,
    email: String,
    sign_in_count: u64,
    active: bool,
}</code></pre>
		
		<h3>Создание экземпляра</h3>
		<pre><code>let user = User {
    email: String::from("user@example.com"),
    username: String::from("user123"),
    active: true,
    sign_in_count: 1,
};</code></pre>
		
		<h3>Методы</h3>
		<pre><code>impl User {
    fn new(email: String, username: String) -> User {
        User {
            email,
            username,
            active: true,
            sign_in_count: 1,
        }
    }
}</code></pre>`,
	},
	{
		Title: "Rust Enums",
		Slug:  "rust-enums",
		Content: `<h2>Перечисления (Enums)</h2>
		
		<h3>Определение enum</h3>
		<pre><code>enum Direction {
    Up,
    Down,
    Left,
    Right,
}

enum IpAddr {
    V4(String),
    V6(String),
}</code></pre>
		
		<h3>Option enum</h3>
		<pre><code>enum Option<T> {
    Some(T),
    None,
}

let some_number = Some(5);
let absent_number: Option<i32> = None;</code></pre>
		
		<h3>match</h3>
		<pre><code>match direction {
    Direction::Up => println!("Вверх"),
    Direction::Down => println!("Вниз"),
    Direction::Left => println!("Влево"),
    Direction::Right => println!("Вправо"),
}</code></pre>`,
	},
}

// C++ Tutorials
var cppTutorials = []TutorialSection{
	{
		Title: "C++ Home",
		Slug:  "cpp-home",
		Content: `<h2>Добро пожаловать в C++!</h2>
		<p>C++ — язык программирования общего назначения, созданный Бьёрном Страуструпом в 1985 году как расширение языка C с классами.</p>
		
		<div style="background: #ffcccb; padding: 20px; border-radius: 12px; margin: 20px 0;">
			<h3>Где используется C++?</h3>
			<ul>
				<li>Игровые движки (Unreal Engine)</li>
				<li>Операционные системы</li>
				<li>Браузеры (Chrome, Firefox)</li>
				<li>Высокопроизводительные вычисления</li>
				<li>Встраиваемые системы</li>
			</ul>
		</div>
		
		<h3>Преимущества C++</h3>
		<ul>
			<li>Высокая производительность</li>
			<li>Полный контроль над памятью</li>
			<li>Мощные возможности ООП</li>
			<li>Огромная экосистема</li>
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
}</code></pre>
		
		<h3>Компиляция и запуск</h3>
		<pre><code>g++ main.cpp -o program
./program</code></pre>`,
	},
	{
		Title: "C++ Syntax",
		Slug:  "cpp-syntax",
		Content: `<h2>Синтаксис C++</h2>
		
		<h3>Переменные</h3>
		<pre><code>int age = 25;
double price = 99.99;
std::string name = "Азамат";
bool isActive = true;
char grade = 'A';</code></pre>
		
		<h3>Ввод и вывод</h3>
		<pre><code>// Вывод
std::cout << "Hello, world!" << std::endl;

// Ввод
int x;
std::cin >> x;</code></pre>`,
	},
	{
		Title: "C++ Functions",
		Slug:  "cpp-functions",
		Content: `<h2>Функции в C++</h2>
		
		<h3>Объявление функций</h3>
		<pre><code>int add(int a, int b) {
    return a + b;
}

void greet(const std::string& name) {
    std::cout << "Привет, " << name << "!" << std::endl;
}</code></pre>
		
		<h3>Прототипы функций</h3>
		<pre><code>// Прототип (объявление)
int multiply(int a, int b);

// Определение
int multiply(int a, int b) {
    return a * b;
}</code></pre>`,
	},
	{
		Title: "C++ Classes",
		Slug:  "cpp-classes",
		Content: `<h2>Классы в C++</h2>
		
		<h3>Определение класса</h3>
		<pre><code>class Person {
private:
    std::string name;
    int age;
    
public:
    Person(std::string name, int age) {
        this->name = name;
        this->age = age;
    }
    
    void greet() {
        std::cout << "Привет, я " << name << std::endl;
    }
};</code></pre>
		
		<h3>Создание объектов</h3>
		<pre><code>Person person("Азамат", 25);
person.greet();</code></pre>`,
	},
	{
		Title: "C++ Pointers",
		Slug:  "cpp-pointers",
		Content: `<h2>Указатели в C++</h2>
		
		<h3>Объявление указателей</h3>
		<pre><code>int x = 42;
int* ptr = &x;  // указатель на x

*ptr = 100;     // изменение через указатель
cout << x;      // 100</code></pre>
		
		<h3>Динамическая память</h3>
		<pre><code>int* arr = new int[10];  // выделение памяти
delete[] arr;             // освобождение памяти</code></pre>`,
	},
	{
		Title: "C++ Vectors",
		Slug:  "cpp-vectors",
		Content: `<h2>Векторы в C++</h2>
		
		<h3>Использование векторов</h3>
		<pre><code>#include <vector>

std::vector<int> v = {1, 2, 3};
v.push_back(4);           // добавить в конец
v.pop_back();             // удалить с конца
int size = v.size();       // размер
bool empty = v.empty();    // проверка на пустоту</code></pre>
		
		<h3>Итерация</h3>
		<pre><code>// По индексу
for (int i = 0; i < v.size(); i++) {
    cout << v[i] << endl;
}

// Range-based for
for (int x : v) {
    cout << x << endl;
}</code></pre>`,
	},
}

// Swift Tutorials
var swiftTutorials = []TutorialSection{
	{
		Title: "Swift Home",
		Slug:  "swift-home",
		Content: `<h2>Добро пожаловать в Swift!</h2>
		<p>Swift — современный язык программирования от Apple, представленный в 2014 году как замена Objective-C для разработки под экосистему Apple.</p>
		
		<div style="background: #ffe4b5; padding: 20px; border-radius: 12px; margin: 20px 0;">
			<h3>Где используется Swift?</h3>
			<ul>
				<li>iOS приложения</li>
				<li>macOS приложения</li>
				<li>watchOS приложения</li>
				<li>tvOS приложения</li>
				<li>Серверная разработка (Vapor)</li>
			</ul>
		</div>
		
		<h3>Преимущества Swift</h3>
		<ul>
			<li>Безопасность (опционалы)</li>
			<li>Современный синтаксис</li>
			<li>Высокая производительность</li>
			<li>Открытый исходный код</li>
		</ul>`,
	},
	{
		Title: "Swift Introduction",
		Slug:  "swift-introduction",
		Content: `<h2>Введение в Swift</h2>
		
		<h3>Первая программа</h3>
		<pre><code>print("Сәлем, Swift!")

let name = "Азамат"
print("Привет, \(name)!")</code></pre>
		
		<h3>Особенности Swift</h3>
		<ul>
			<li>Type Safety и Type Inference</li>
			<li>Опционалы для безопасной работы с nil</li>
			<li>Функциональное программирование</li>
			<li>Protocol-Oriented Programming</li>
		</ul>`,
	},
	{
		Title: "Swift Variables",
		Slug:  "swift-variables",
		Content: `<h2>Переменные в Swift</h2>
		
		<h3>var и let</h3>
		<pre><code>var age = 25        // изменяемая
age = 26            // OK

let name = "Азамат"  // константа
// name = "Диас"    // ошибка!</code></pre>
		
		<h3>Типы данных</h3>
		<pre><code>let age: Int = 25
let price: Double = 99.99
let name: String = "Азамат"
let isActive: Bool = true</code></pre>`,
	},
	{
		Title: "Swift Optionals",
		Slug:  "swift-optionals",
		Content: `<h2>Опциональные типы</h2>
		
		<h3>Объявление опционалов</h3>
		<pre><code>var age: Int? = 25
var name: String? = nil</code></pre>
		
		<h3>Извлечение опционалов</h3>
		<pre><code>// if let
if let age = age {
    print("Возраст: \(age)")
}

// guard let
guard let age = age else { return }

// force unwrap (опасно!)
let value = age!</code></pre>
		
		<h3>Опциональная цепочка</h3>
		<pre><code>let city = user.address?.city</code></pre>`,
	},
	{
		Title: "Swift Functions",
		Slug:  "swift-functions",
		Content: `<h2>Функции в Swift</h2>
		
		<h3>Объявление функций</h3>
		<pre><code>func greet(name: String) -> String {
    return "Привет, \(name)!"
}

func add(_ a: Int, _ b: Int) -> Int {
    return a + b
}</code></pre>
		
		<h3>Параметры по умолчанию</h3>
		<pre><code>func greet(name: String = "гость") -> String {
    return "Привет, \(name)!"
}</code></pre>`,
	},
	{
		Title: "Swift Classes",
		Slug:  "swift-classes",
		Content: `<h2>Классы в Swift</h2>
		
		<h3>Определение класса</h3>
		<pre><code>class Person {
    var name: String
    var age: Int
    
    init(name: String, age: Int) {
        self.name = name
        self.age = age
    }
    
    func greet() -> String {
        return "Привет, я \(name)"
    }
}</code></pre>
		
		<h3>Наследование</h3>
		<pre><code>class Student: Person {
    var university: String
    
    init(name: String, age: Int, university: String) {
        self.university = university
        super.init(name: name, age: age)
    }
    
    override func greet() -> String {
        return "Я студент \(university)"
    }
}</code></pre>`,
	},
	{
		Title: "Swift Structs",
		Slug:  "swift-structs",
		Content: `<h2>Структуры в Swift</h2>
		
		<h3>Определение структуры</h3>
		<pre><code>struct Point {
    var x: Double
    var y: Double
    
    func distance(to point: Point) -> Double {
        let dx = x - point.x
        let dy = y - point.y
        return sqrt(dx*dx + dy*dy)
    }
}</code></pre>
		
		<h3>Классы vs Структуры</h3>
		<ul>
			<li>Классы — ссылочный тип</li>
			<li>Структуры — значимый тип</li>
			<li>Структуры не поддерживают наследование</li>
			<li>Структуры имеют автоматический инициализатор</li>
		</ul>`,
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

	// Основные страницы
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/top", TopLanguagesHandler)
	http.HandleFunc("/search", SearchHandler)

	// Новые обработчики для фильтрации
	http.HandleFunc("/filter", FilterHandler)
	http.HandleFunc("/api/filter", FilterAPIHandler)
	http.HandleFunc("/category/", CategoryHandler)
	http.HandleFunc("/difficulty/", DifficultyHandler)

	// Аутентификация
	http.HandleFunc("/register", RegisterHandler)
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/logout", LogoutHandler)

	// Детали языка и профиль
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

// ГЛАВНАЯ СТРАНИЦА С ФИЛЬТРАЦИЕЙ
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	username := getUsernameFromCookie(r)

	// Получаем параметры фильтрации из URL
	category := r.URL.Query().Get("category")
	difficulty := r.URL.Query().Get("difficulty")
	sortBy := r.URL.Query().Get("sort")
	search := r.URL.Query().Get("q")

	if sortBy == "" {
		sortBy = "popularity"
	}

	// Получаем языки с фильтрацией
	var languages []*models.Language

	if search != "" {
		// Поиск
		languages = storage.SearchLanguages(search)
	} else if category != "" || difficulty != "" {
		// Фильтрация
		filters := make(map[string]interface{})
		if category != "" {
			filters["category"] = category
		}
		if difficulty != "" {
			filters["difficulty"] = difficulty
		}
		languages = storage.FilterLanguages(filters)
	} else {
		// Обычная сортировка
		languages = storage.GetAllLanguages("", sortBy)
	}

	// Получаем все категории для фильтра
	categories := storage.GetAllCategories()

	data := struct {
		Username     string
		Languages    []*models.Language
		Categories   []*models.Category
		SelectedCat  string
		SelectedDiff string
		CurrentSort  string
		SearchQuery  string
		CurrentPage  string
	}{
		Username:     username,
		Languages:    languages,
		Categories:   categories,
		SelectedCat:  category,
		SelectedDiff: difficulty,
		CurrentSort:  sortBy,
		SearchQuery:  search,
		CurrentPage:  "home",
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Ошибка шаблона: "+err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
	tmpl.Execute(w, data)
}

// Страница фильтрации
func FilterHandler(w http.ResponseWriter, r *http.Request) {
	username := getUsernameFromCookie(r)

	category := r.URL.Query().Get("category")
	difficulty := r.URL.Query().Get("difficulty")
	sortBy := r.URL.Query().Get("sort")

	if sortBy == "" {
		sortBy = "popularity"
	}

	filters := make(map[string]interface{})
	if category != "" && category != "all" {
		filters["category"] = category
	}
	if difficulty != "" && difficulty != "all" {
		filters["difficulty"] = difficulty
	}

	var languages []*models.Language
	if len(filters) > 0 {
		languages = storage.FilterLanguages(filters)
	} else {
		languages = storage.GetAllLanguages("", sortBy)
	}

	categories := storage.GetAllCategories()

	data := struct {
		Username     string
		Languages    []*models.Language
		Categories   []*models.Category
		SelectedCat  string
		SelectedDiff string
		CurrentSort  string
		Query        string
	}{
		Username:     username,
		Languages:    languages,
		Categories:   categories,
		SelectedCat:  category,
		SelectedDiff: difficulty,
		CurrentSort:  sortBy,
		Query:        "",
	}

	tmpl, err := template.ParseFiles("templates/filter.html")
	if err != nil {
		http.Error(w, "Ошибка шаблона: "+err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}

// API для AJAX фильтрации
func FilterAPIHandler(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")
	difficulty := r.URL.Query().Get("difficulty")

	filters := make(map[string]interface{})
	if category != "" && category != "all" {
		filters["category"] = category
	}
	if difficulty != "" && difficulty != "all" {
		filters["difficulty"] = difficulty
	}

	var languages []*models.Language
	if len(filters) > 0 {
		languages = storage.FilterLanguages(filters)
	} else {
		languages = storage.GetAllLanguages("", "popularity")
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(languages)
}

// Страница категории
func CategoryHandler(w http.ResponseWriter, r *http.Request) {
	category := strings.TrimPrefix(r.URL.Path, "/category/")
	username := getUsernameFromCookie(r)

	filters := map[string]interface{}{"category": category}
	languages := storage.FilterLanguages(filters)

	data := struct {
		Username  string
		Languages []*models.Language
		Category  string
		Query     string
	}{
		Username:  username,
		Languages: languages,
		Category:  category,
		Query:     "",
	}

	tmpl, err := template.ParseFiles("templates/category.html")
	if err != nil {
		http.Error(w, "Ошибка шаблона", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}

// Страница по сложности
func DifficultyHandler(w http.ResponseWriter, r *http.Request) {
	difficulty := strings.TrimPrefix(r.URL.Path, "/difficulty/")
	username := getUsernameFromCookie(r)

	filters := map[string]interface{}{"difficulty": difficulty}
	languages := storage.FilterLanguages(filters)

	data := struct {
		Username   string
		Languages  []*models.Language
		Difficulty string
		Query      string
	}{
		Username:   username,
		Languages:  languages,
		Difficulty: difficulty,
		Query:      "",
	}

	tmpl, err := template.ParseFiles("templates/difficulty.html")
	if err != nil {
		http.Error(w, "Ошибка шаблона", http.StatusInternalServerError)
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
	categories := storage.GetAllCategories()

	data := struct {
		Username    string
		Languages   []*models.Language
		Categories  []*models.Category
		Query       string
		SelectedCat string
	}{
		Username:    username,
		Languages:   languages,
		Categories:  categories,
		Query:       query,
		SelectedCat: "",
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
	sectionSlug := "home"
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

	// Увеличиваем счетчик просмотров
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

// ТОП-20 языков программирования
func TopLanguagesHandler(w http.ResponseWriter, r *http.Request) {
	username := getUsernameFromCookie(r)

	data := struct {
		Username    string
		CurrentPage string
	}{
		Username:    username,
		CurrentPage: "top",
	}

	tmpl, err := template.ParseFiles("templates/top.html")
	if err != nil {
		http.Error(w, "Ошибка шаблона: "+err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}
