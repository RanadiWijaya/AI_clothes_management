# 👗 AI Fashion Stylist (CLI)

A simple command-line interface (CLI) application written in Go that helps users manage their clothing inventory and get outfit recommendations based on category and formality.

## ✨ Features

- ➕ Add clothing items with category, color, and formality level
- 📄 View all clothing in the wardrobe
- 🔍 Search clothing by color or category
- 🔃 Sort clothing by formality (Selection Sort)
- 🤖 Generate outfit recommendations (top + bottom)

## 🧵 Categories

-  atas – top wear (e.g., shirts, blouses)
-  bawah – bottom wear (e.g., pants, skirts)
-  luar – outerwear (e.g., jackets, coats)
-  aksesoris – accessories (e.g., belts, watches)

## 🧮 Algorithms Used

### 1. **Selection Sort** 🔄
- **Purpose**: Sort clothing items by formality level (1-5)
- **Implementation**: selectionSort() function
- **How it works**: 
  - Finds the minimum element and swaps it with the first element
  - Repeats for the remaining unsorted portion
  - Sorts in ascending order of formality

### 2. **Sequential Search (Linear Search)** 🔍
- **Purpose**: Search clothing items by color or category
- **Implementation**: pencarianPakaian() function
- **How it works**:
  - Iterates through each item in the wardrobe
  - Compares search criteria with item properties
  - Returns all matching items

### 3. **Simple Matching Algorithm** 🤖
- **Purpose**: Generate outfit recommendations by pairing tops with bottoms
- **Implementation**: pisahkanKategoriPakaian() and tampilkanRekomendasi() functions
- **How it works**:
  - Separates clothing into categories (tops vs bottoms)
  - Creates combinations by pairing items sequentially
  - Limits recommendations to maximum of 3 outfits

## 📦 Installation

Make sure you have Go installed on your system. Then:

bash
git clone https://github.com/your-username/ai-fashion-stylist.git
cd ai-fashion-stylist
go run main.go

## 🧑‍💻 Usage

=== AI Fashion Stylist ===
1. Tambah Pakaian
2. Lihat Daftar Pakaian
3. Cari Pakaian
4. Urutkan Pakaian
5. Rekomendasi Outfit
0. Keluar


## 🔧 Tech Stack

- **Go (Golang)** – Core language for building the CLI app
- **Standard libraries only** (no third-party packages)
- **Built-in algorithms** for sorting and searching operations

## 🛠 Sample Workflow

1. Add a few clothing items (top and bottom)
2. Use the "Rekomendasi Outfit" menu to get suggested outfit combinations
3. Sort your wardrobe based on formality to organize looks for events
4. Search for specific items by color or category when planning outfits

## 📊 Data Structure

The application uses a fixed-size array ([MAX_PAKAIAN]Pakaian) to store up to 100 clothing items, where each item contains:

go
type Pakaian struct {
    Tipe              string  // Type of clothing
    Kategori          string  // Category (atas/bawah/luar/aksesoris)
    Warna             string  // Color
    tingkatFormalitas int     // Formality level (1-5)
}


## 🎯 Algorithm Design Choices

- **Selection Sort**: Chosen for its simplicity and in-place sorting capability, suitable for small datasets
- **Sequential Search**: Implemented to handle unsorted data and find multiple matches
- **Array-based Storage**: Uses fixed arrays for predictable memory usage and simple indexing
