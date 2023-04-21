# Secure Pass

Secure Pass เป็นเครื่องมือเข้ารหัสและถอดรหัสด้วยการเข้ารหัส AES-256-CBC ด้วยคีย์ขนาด 32 ไบต์และเวกเตอร์เริ่มต้นขนาด 16 ไบต์ โปรแกรมสามารถรับรหัสผ่านข้อความธรรมดาหรือรหัสผ่านที่ถูกเข้ารหัสเป็น base64 และสามารถส่งออกเป็นรหัสผ่านที่ถูกเข้ารหัสเป็น base64 หรือรหัสผ่านข้อความธรรมดาตามลำดับ

## วิธีการใช้งาน

[ดาวน์โหลดได้ที่หน้า release page ของ Secure Pass](https://github.com/wisesight/secure-pass/releases) ตามระบบปฏิบัติการและโครงสร้างของคุณ

```
Usage: ./secure-pass (-e|--encrypt) plaintext OR (-d|--decrypt) ciphertext
```

โปรแกรมรับอินพุตเป็นพารามิเตอร์ ดังนี้

- `-e` หรือ `--encrypt`: ทำการเข้ารหัสผ่านแบบข้อความธรรมดา และส่งออกเป็นรหัสผ่านที่ถูกเข้ารหัสเป็น base64
- `-d` หรือ `--decrypt`: ทำการถอดรหัสผ่านที่ถูกเข้ารหัสเป็น base64 และส่งออกเป็นรหัสผ่านข้อความธรรมดา

ตัวอย่าง:

การเข้ารหัสผ่าน:

```bash
./secure-pass -e mypassword
```

การถอดรหัสผ่าน:

```bash
./secure-pass -d YnJvd3NlcjpteXBhc3N3b3Jk
```

## วิธีการสร้าง

**ถ้าคุณต้องการสร้างด้วยตัวเอง**

### ข้อกำหนด

เพื่อใช้ Secure Pass คุณต้องติดตั้งรายการดังนี้บนเครื่องของคุณ:

- Go 1.16 หรือใหม่กว่า
- make

1. Clone project หรือดาวน์โหลดเป็นไฟล์ ZIP:

```bash
git clone https://github.com/example/secure-pass.git
```

2. สร้างไฟล์ executable:

```bash
cd secure-pass
make --always-make build VERSION=X.X.X
```

1. โปรดตรวจสอบไฟล์ executable ใน build directory 

## การเรียกใช้ใน NodeJs

โปรดตรวจสอบใน example directory

---

# Secure Pass

Secure Pass is a command-line tool for encrypting and decrypting passwords using AES-256-CBC encryption with a 32-byte key and 16-byte initialization vector (IV). The tool takes in plaintext passwords or base64-encoded ciphertext passwords and outputs ciphertext passwords or plaintext passwords, respectively.

## Usage

[Download it from the release page](https://github.com/wisesight/secure-pass/releases) according to your operating system and architecture.

```
Usage: ./secure-pass (-e|--encrypt) plaintext OR (-d|--decrypt) ciphertext
```

The tool takes in two arguments:

- `-e` or `--encrypt`: encrypts the plaintext password and outputs the base64-encoded ciphertext password
- `-d` or `--decrypt`: decrypts the base64-encoded ciphertext password and outputs the plaintext password

Examples:

Encrypting a password:

```bash
./secure-pass -e mypassword
```

Decrypting a password:

```bash
./secure-pass -d YnJvd3NlcjpteXBhc3N3b3Jk
```

## Build

**if you want to build it by yourself.**

### Requirements

To use Secure Pass, you need to have the following installed on your machine:

- Go 1.16 or later
- make

1. Clone the repository or download it as a ZIP file:

```bash
git clone https://github.com/example/secure-pass.git
```

2. Build the executable:

```bash
cd secure-pass
make --always-make build VERSION=X.X.X
```

3. Please check the executable files in the build directory.

## Execute a binary file in NodeJs

Please check in the example directory