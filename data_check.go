package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// メイン処理
func main() {
	path_txt := "path.txt"
	path := read_path(path_txt)

	file_datas := search_files(path)

	ratio_data := ratio(file_datas)

	selected_data := select_content(ratio_data)
	fmt.Println(selected_data)
}

// パス読み込み
func read_path(path_data string) string {
	data, err := ioutil.ReadFile(path_data)
	if err != nil {
		panic(err)
	}

	path := string(data)

	return path
}

// ファイルデータ取得
func search_files(search_path string) map[string]int {
	file_datas, err := ioutil.ReadDir(search_path)
	if err != nil {
		panic(err)
	}

	// ファイル用変数定義
	image := 0
	movie := 0
	music := 0
	program := 0
	text := 0
	office := 0

	for _, file_data := range file_datas {
		fmt.Printf("%s: %d\n", file_data.Name(), file_data.Size())
		ext_name := filepath.Ext(file_data.Name())
		fmt.Println(ext_name)
		// 拡張子別にデータ量加算
		if ext_name == ".png" || ext_name == ".jpg" {
			image += int(file_data.Size())
		} else if ext_name == ".mp4" || ext_name == ".avi" {
			movie += int(file_data.Size())
		} else if ext_name == ".mp3" {
			music += int(file_data.Size())
		} else if ext_name == ".py" || ext_name == ".go" {
			program += int(file_data.Size())
		} else if ext_name == ".txt" || ext_name == ".pdf" || ext_name == ".csv" {
			text += int(file_data.Size())
		} else if ext_name == ".docx" || ext_name == ".xlsx" {
			office += int(file_data.Size())
		}
	}
	dict := map[string]int{
		"画像":       image,
		"映像":       movie,
		"音楽":       music,
		"プログラム":    program,
		"テキストファイル": text,
		"オフィスファイル": office,
	}

	return dict
}

// 比率を求める
func ratio(dict_data map[string]int) map[string]float64 {
	sum := 0.0
	for _, value := range dict_data {
		sum += float64(value)
	}

	image_value := float64(dict_data["画像"])
	movie_value := float64(dict_data["映像"])
	music_value := float64(dict_data["音楽"])
	program_value := float64(dict_data["プログラム"])
	text_value := float64(dict_data["テキストファイル"])
	office_value := float64(dict_data["オフィスファイル"])

	new_dict := map[string]float64{
		"画像":       image_value / sum * 100,
		"映像":       movie_value / sum * 100,
		"音楽":       music_value / sum * 100,
		"プログラム":    program_value / sum * 100,
		"テキストファイル": text_value / sum * 100,
		"オフィスファイル": office_value / sum * 100,
	}

	fmt.Println(new_dict)

	return new_dict
}

// 0%の情報を除外
func select_content(ratio_data map[string]float64) map[string]float64 {

	// 0%を除外した辞書定義
	select_dict := map[string]float64{}

	for key, value := range ratio_data {
		if value != 0 {
			select_dict[key] = ratio_data[key]
		}
	}

	fmt.Println(select_dict)

	return select_dict
}

func write_text(selected_data map[string]float64) {
	text_file, err := os.Create("log.txt")
	if err != nil {
		panic(err)
	}

	defer text_file.Close()

	for key, value := range selected_data {
		text_file.WriteString(key + ":" + strconv.FormatFloat(value, 'f', 2, 64) + "%\n")
	}
}
