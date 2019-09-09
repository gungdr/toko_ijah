# Toko Ijah

This application is a REST API to handling inventory for a small business

## Story

Ijah usually notes her inventory activity in a spreadsheet document.
Ijah wants to replace her spreadsheet by creating an application.

### Requirements

Must have features :

1. Catatan Jumlah Barang [X]
   * Stores actual stock of products
2. Catatan Barang Masuk [X]
   * Catatan Barang Masuk
3. Catatan Barang Keluar [X]
   * To store product that will be stored into the inventory.
4. Laporan Nilai Barang [X]
   * Shows a report for ijah to help her analyze and make decision. This report is related to total inventory value of Toko Ijah.
5. Laporan Penjualan [X]
   * Shows a report for ijah to help her analyze and make decision. This report is related to omzet / selling / profit.
6. Export data report in CSV format [X]

Optional features:

1. import data from CSV/spreadsheet (data migration) []
2. CMS UI for inventory management []

### Assumptions

1. The **SKU** is a combination of product attributes such as store code `SSI-`,brand code `DXXXXX-`, size code `MM-`, and color code `RED`
2. **ID pesanan** is combination of sales prefix code `ID-`, date `YYYYMMDD-`, and running number `531154`
3. There are processes such as Purchase order, Receive order, and Sales.
4. The **Nomer Kwitansi** is not mandatory when receiving purchase order.
5. In Stock product is available after the Receipt order process.
6. Purchase order can be completed with one to several receipt order.
7. Out stock will be counted from sales

## Design

Base on requirements, the structure will be look like this.

![toko ijah](doc/inventory.svg)

## Stack

This application is use golang and sqlite.

## Project structure

this project has 4 main structure

```bash
models
repository
usecase
delivery
```

### models

Models is where all the entities are

### repository

Repository connecting app to external service such as `DB` or `REST API`

### usecase

Usecase is place for all business rules

### delivery

Delivery is how we present our data.
we can serve it as `REST API`, `web`, `GraphQL` etc.

## Running the executable binary

In Linux

```bash
./start ./toko_ijah_linux
```

In Mac

```bash
./start ./toko_ijah_mac
```

## Available Scripts

```bash
make run
```

this script is retrieving dependencies, build and run the application

```bash
make test
```

this script run the test

```bash
make dev
```

this script run the app in dev mode, the app runs without building

```bash
make clean
```

this script clean the executable

## Endpoints

> GET    /report/in-stock  (Catatan Barang Masuk)
>
> GET    /report/out-stock (Catatan Barang Keluar)
>
> GET    /report/total-product (Catatan Jumlah Barang)
>
> GET    /report/product-value (Laporan Nilai Barang)
>
> GET    /report/sales?from={YYYY-mm-dd}&until={YYYY-mm-dd} (Laporan Penjualan)

All endpoints return json.
To export data as CSV append `?format=csv` at the end of query string.

```REST
/report/in-stock?format=csv

/report/sales?from=2019-01-01&until=2018-02-01&format=csv
```

## TODO

this application is not having a full features, these are the features need to be implemented

1. Implement purchase order transaction process
2. Implement receive order transaction process
3. Implement stock opname process
4. Implement sales transaction process
5. Implement data migration
6. CMS UI
7. Create Docker images so it can run anywhere
