# **Sales Data Analysis System**  

A high-performance sales data analysis system built in **Go**, designed to efficiently process and analyze large CSV files containing sales data.  

## 🚀 **Features**  

- **Large CSV File Processing**: Efficient handling of large sales datasets  
- **Data Analysis**: Comprehensive sales data insights and reporting  
- **Performance Optimized**: Built for speed and memory efficiency  
- **Scalable Architecture**: Designed to support growing data volumes  



## 📋 **Prerequisites**  

Ensure you have the following installed:  

- **Go**: Version 1.19 or higher  
- **Operating System**: Linux, macOS, or Windows  



## 🛠️ **Installation & Setup**  

1. **Clone the repository:**  
   ```bash
   git clone git@github.com:naveenraja2310/lumel.git
   cd lumel
   ```

2. **Install dependencies:**  
   ```bash
   go mod tidy
   ```

3. **Run the application:**  
   ```bash
   go run main.go
   ```



## 🔧 **Usage**  

### **Basic Usage**  
```bash
# Start the application
go run main.go
```

### **Configuration**  
Place your CSV files in the root directory with the filename **sample.csv**.  



## 📈 **Performance**  

- **Memory Efficient**: Uses streaming processing to handle files larger than available RAM  
- **Fast Processing**: Optimized algorithms for quick and accurate analysis  
- **Concurrent Processing**: Leverages Go's concurrency features for parallel execution  

## 🏗️ **Schema Diagram**  
For a visual representation of the database structure, refer to **schema.png** in the root folder.

## ⏳ **Data Processing via CRON Job**  
The **CRON job runs every midnight**, updating existing data if available or creating new records if absent.  

## 🧪 **API Testing with Postman**  

This project includes a Postman collection for easy API testing and development.  

### **How to Import the Postman Collection:**  

1. Open the **Postman Application** (Download and install if needed).  
2. Click the **Import** button in the top-left corner of Postman.  
3. Select the **Upload Files** tab.  
4. Click **Choose Files** and select the provided `.json` collection file.  
5. Alternatively, drag and drop the collection file directly into Postman.  

### **Postman Collection File:**  
You can find **Lumel.postman_collection.json** in the root directory.  



## 🌐 **API Endpoints**  

### **Data Loading**  
- `GET /load-data` - Load data from CSV file (CRON job runs every midnight)  

### **Analytics**  
- `GET /analytics/total-revenue?start=2020-12-01&end=2025-12-31` - Get total revenue  
- `GET /analytics/total-revenue-by-category?start=2020-12-01&end=2025-12-31` - Get total revenue by category  
- `GET /analytics/total-revenue-by-region?start=2020-12-01&end=2025-12-31` - Get total revenue by region  
- `GET /analytics/total-revenue-by-product?start=2020-12-01&end=2025-12-31` - Get total revenue by product  



## ⚙️ **Environment Variables**  

Configure the following environment variables before running the application:  

```bash
DB_URI="mongodb+srv://<username>:<password>@cluster0.x6sdhax.mongodb.net/"
DB_NAME="lumel"
DB_TIME="10"
APP_PORT="9000"
```

