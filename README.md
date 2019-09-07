# Toko Ijah

This application is a REST API to handling inventory for a small business

## Story

Ijah usually notes her inventory activity in a spreadsheet document.
Ijah wants to replace her spreadsheet by creating an application.

### Requirements

Must have features :

1. Catatan Jumlah Barang
   * Stores actual stock of products
2. Catatan Barang Masuk
   * Catatan Barang Masuk
3. Catatan Barang Keluar
   * To store product that will be stored into the inventory.
4. Laporan Nilai Barang
   * Shows a report for ijah to help her analyze and make decision. This report is related to total inventory value of Toko Ijah.
5. Laporan Penjualan
   * Shows a report for ijah to help her analyze and make decision. This report is related to omzet / selling / profit.
6. Export data report in CSV format

Optionnal features:

1. import data from CSV/spreadsheet (data migration)
2. CMS UI for inventory management

## Stack

This application is use golang and sqlite.

### Run the application

```bash
make run
```
