package main

import (
	"fmt"
	"math"
)

func main() {
	// typeを使って基本型にエイリアスを貼ることができる
	type MyInt int
	var n1 MyInt
	n2 := MyInt(7)
	fmt.Println(n1, n2)

	// typeにより様々なエイリアスを貼ることができる
	type (
		IntPair     [2]int
		Strings     []string
		AreaMap     map[string][2]float64
		IntsChannel chan []int
	)
	// エイリアスを設定することにより、リテラルの複雑な記述を簡略化でき、プログラムの見通しも良くなる
	pair := IntPair{1, 2}
	strs := Strings{"Apple", "Banana", "Cherry"}
	amap := AreaMap{
		"Tokyo": {4.12, 3.15},
		"Osaka": {3.52, 4.25},
	}
	ich := make(IntsChannel)
	fmt.Println(pair, strs, amap, ich)

	// 関数の型のエイリアスも作成できる
	type Callback func(i int) int

	sum := func(ints []int, callback Callback) int {
		var sum int
		for _, v := range ints {
			sum += v
		}

		return callback(sum)
	}([]int{1, 2, 3, 4, 5}, func(i int) int {
		return i * 2
	})
	fmt.Println(sum)

	type T0 int
	type T1 int
	// エイリアスと基本型には互換性がある
	t0 := T0(5)
	i0 := int(t0)
	fmt.Println(t0, i0)
	// こちらも同様
	t1 := T1(8)
	i1 := int(t1)
	fmt.Println(t1, i1)
	// コンパイルエラー、エイリアス間には互換性がない
	//t0 = t1

	// 構造体、値渡しとなる
	type Point struct {
		X int
		Y int
		// こちらでも同じ意味
		// X, Y int
	}

	var pt Point
	fmt.Println(pt.X, pt.Y)
	pt.X = 10
	pt.Y = 8
	fmt.Println(pt.X, pt.Y)

	// 構造体の定義された順番に値を入れることでプロパティに値をセットすることができる
	pt2 := Point{1, 3}
	fmt.Println(pt2)
	// プロパティ名を明示的に指定して定義することもでき、こちらのほうが視覚的に得られる情報は多い
	pt2 = Point{
		X: 23,
		Y: 26,
	}
	// プロパティを一部だけ初期化する事もできる
	fmt.Println(pt2)
	pt2 = Point{
		Y: 3452,
	}
	fmt.Println(pt2)

	// プロパティ名が型名と同じで良ければ、プロパティ名を省略できる
	type T struct {
		int
		string
		float64
	}
	pt3 := T{3, "test", 3.2}
	fmt.Println(pt3, pt3.int, pt3.string, pt3.float64)

	// 構造体の中に構造体を定義することもできる
	type Feed struct {
		Name   string
		Amount int
	}
	type Animal struct {
		Name string
		Feed Feed
		// 型名と同じプロパティ名で良ければ以下のような定義でもOK
		// Feed
		// 基本型でプロパティ名を省略することはケースとしてあまりありえないが
		// 構造体の中に構造体を定義するのであれば、積極的に使うことで読みやすいコードになる
	}
	a := Animal{
		Name: "Monkey",
		Feed: Feed{
			Name:   "Banana",
			Amount: 10,
		},
	}
	fmt.Println(a, a.Feed.Name)

	type NestT0 struct {
		Name1 string
		name  string
	}

	type NestT1 struct {
		NestT0
		Name2 string
		name  string
	}

	type NestT2 struct {
		NestT1
		Name3 string
		name  string
	}

	// 構造体をネストさせて、値を入れる
	t2 := NestT2{
		NestT1{
			NestT0{
				Name1: "Name1",
				name:  "name 1",
			},
			"Name2",
			"name 2",
		},
		"Name3",
		"name 3",
	}
	// ネストしたとしてもプロパティ名が一意に定まる場合はプロパティ名を省略できる
	// 一意に定まらない場合は中間のプロパティに対してアクセスすることで値を取り出せる
	fmt.Println(t2.Name1, t2.Name2, t2.Name3, t2.name, t2.NestT1.name, t2.NestT1.NestT0.name)

	// この性質を生かして共通化ができる
	type Base struct {
		Id    int
		Owner string
	}

	type A struct {
		Base
		Name string
		Area string
	}

	type B struct {
		Base
		Title  string
		Bodies []string
	}

	a1 := A{
		Base: Base{
			Id:    11,
			Owner: "Yasu",
		},
		Name: "Taro",
		Area: "Tokyo",
	}
	b := B{
		Base: Base{
			Id:    12,
			Owner: "Hanako",
		},
		Title:  "no title",
		Bodies: []string{"A", "B"},
	}
	fmt.Println(a1, b, a1.Id, a1.Owner, b.Id, b.Owner)

	type IncPointer struct {
		T
		*T1 // 構造体のポインタ型を定義するとき、フィールド名はT1となる
	}

	type Point2 struct {
		x, y int
	}
	// 構造体のポインタを受け取る
	swap := func(p2 *Point2) {
		// s[i]のようにデリファレンスなしでアクセスできる
		x, y := p2.x, p2.y

		p2.x = y
		p2.y = x
	}
	// 構造体から直接ポインタを取り出すパターンはGoでも頻出
	p := &Point2{
		x: 12312,
		y: 9809,
	}
	// 構造体のポインタを渡す
	swap(p)
	fmt.Println(p.x, p.y)

	type Person struct {
		Id   int
		Name string
		Area string
	}
	// newでポインタ型を生成することができる
	// newは構造体専用というわけではなくnew(int)というふうにできるがメリットがなく、構造体のポインタ生成のために使われることが多い
	person := new(Person)
	// person := & Person{ Id: 1, Name: "name", Area: "area" }
	// こういうふうに書くのと大差ないため状況によって使い分ける
	fmt.Println(person.Id, person.Name, person.Area)

	p2 := &Point3{
		X: 3,
		Y: 15,
	}
	p2.Render()

	dp := &Point3{
		X: 234,
		Y: 345,
	}
	fmt.Println(p2.Distance(dp))

	fmt.Println(MyInt2(3).Plus(7))

	ip := IntPair2{1, 0}
	fmt.Println(ip.First(), ip.Last())

	fmt.Println(Strings2{"A", "B", "C"}.Join(","))

	fmt.Println(NewUser(1, "Taro"))

	// 構造体のメソッドを変数にわたすこともできる
	f := (*Point10).ToString
	fmt.Println(f(&Point10{X: 8, Y: 11}))
	// 短縮するとこうなる
	fmt.Println(((*Point10).ToString)(&Point10{
		X: 23,
		Y: 34,
	}))
	// このように書くこともできる
	p3 := &Point10{
		X: 22,
		Y: 33,
	}
	f2 := p3.ToString
	fmt.Println(f2())

	ps := make([]Point20, 5)
	for _, v := range ps {
		fmt.Println(v.X, v.Y)
	}

	pts := Points{}
	pts = append(pts, &Point20{X: 3, Y: 8})
	pts = append(pts, nil)
	pts = append(pts, &Point20{X: 3, Y: 4})
	fmt.Println(pts.ToString())
}

type Point3 struct {
	X, Y int
}

// 構造体のメソッド
// メソッドをはやしたい構造体のポインタを引数に指定する
func (p *Point3) Render() {
	// 構造体の値にアクセスできる
	fmt.Printf("<%d, %d>\n", p.X, p.Y)
}

// 複数メソッドを生やすこともできる
func (p *Point3) Distance(dp *Point3) float64 {
	x, y := p.X-dp.X, p.Y-dp.Y
	return math.Sqrt(float64(x*x + y*y))
}

// 型のエイリアスに対してメソッドを実装することもできる
type MyInt2 int

func (m MyInt2) Plus(i int) int {
	return int(m) + i
}

type IntPair2 [2]int

func (ip IntPair2) First() int {
	return ip[0]
}
func (ip IntPair2) Last() int {
	return ip[1]
}

type Strings2 []string

func (s Strings2) Join(d string) string {
	sum := ""
	for _, v := range s {
		if sum != "" {
			sum += d
		}
		sum += v
	}

	return sum
}

// Goにはconstructorの機能は存在しないが、慣例的に実装することが多い
type User struct {
	// フィールドにタグというメタ情報を付与することができる
	// reflectパッケージを使うとこのタグにアクセスできる
	// この例でいうとjsonパッケージを使ってJSONのキーをタグから設定できる
	Id   int    `json:"user_id"`
	Name string `json:"name"`
}

// 構造体を生成するメソッドとしてNewXXXというメソッドを実装することが頻出らしい
func NewUser(id int, name string) *User {
	u := new(User)
	u.Id = id
	u.Name = name

	return u
}

// パッケージ外部に公開できる構造体
type Point10 struct {
	// 公開できるフィールド
	X, Y int
	// 公開できないフィールド
	z int
}

// 公開できるメソッド
func (p *Point10) ToString() string {
	return fmt.Sprintf("[%d, %d]", p.X, p.Y)
}

// 公開できないメソッド
func (p *Point10) toString() string {
	return fmt.Sprintf("[%d, %d]", p.X, p.Y)
}

// スライスと構造体を組み合わせて使うことが多い
type Point20 struct {
	X, Y int
}

// このようにして組み合わせる
type Points []*Point20

// 組み合わせたエイリアスにメソッドを実装することでデータを扱いやすくすることが重要
func (ps Points) ToString() string {
	str := ""
	for _, p := range ps {
		if str != "" {
			str += ","
		}
		if p == nil {
			str += "<nil>"
		} else {
			str += fmt.Sprintf("[%d, %d]", p.X, p.Y)
		}
	}

	return str
}
