# Docker Configuration

Hướng dẫn cấu hình và sử dụng Docker cho dự án Company Management.

## Cấu trúc thư mục

```
docker/
├── docker-compose.yml         # Cấu hình cho môi trường development
├── docker-compose.prod.yml    # Cấu hình cho môi trường production
├── mysql/
│   ├── init/                  # Scripts khởi tạo MySQL
│   └── backup/                # Thư mục chứa backup database
└── README.md                  # File này
```

## Môi trường Development

### Bắt đầu sử dụng

```bash
# Từ thư mục root của project
make dev
```

Hoặc sử dụng docker-compose trực tiếp:

```bash
cd docker
docker compose -f docker-compose.yml up -d
```

### Thông tin kết nối

- **Server API**: http://localhost:8080
- **Nginx**: http://localhost
- **MySQL**: localhost:3306
  - Database: `company_db`
  - User: `dev_user`
  - Password: `dev_password`

### Tính năng

- Hot-reload với Air
- MySQL 8.0
- Nginx reverse proxy
- Volume mapping cho development

## Môi trường Production

### Cấu hình biến môi trường

Tạo file `.env` trong thư mục `docker/`:

```env
# MySQL Configuration
MYSQL_ROOT_PASSWORD=your_secure_root_password
MYSQL_DATABASE=company_db
MYSQL_USER=prod_user
MYSQL_PASSWORD=your_secure_password

# Application Configuration
JWT_SECRET=your_jwt_secret_key_here
API_KEY=your_api_key_here

# Timezone
TZ=Asia/Ho_Chi_Minh
```

### Cấu hình SSL

Đặt SSL certificates vào thư mục `nginx/ssl/`:

```bash
nginx/ssl/
├── cert.pem
└── key.pem
```

### Bắt đầu sử dụng

```bash
# Từ thư mục root của project
make prod
```

Hoặc sử dụng docker-compose trực tiếp:

```bash
cd docker
docker compose -f docker-compose.prod.yml up -d
```

### Thông tin kết nối

- **Server API**: http://localhost:8080
- **Nginx HTTP**: http://localhost (tự động redirect sang HTTPS)
- **Nginx HTTPS**: https://localhost
- **MySQL**: localhost:3306

### Tính năng

- Optimized build với multi-stage
- SSL/TLS support
- Rate limiting
- Caching
- Security headers
- Health checks
- Resource limits

## Makefile Commands

### Development

```bash
make dev            # Khởi động môi trường development
make build-dev      # Build images cho development
make up-dev         # Start containers
make down-dev       # Stop containers
make restart-dev    # Restart containers
make logs-dev       # Xem logs
make db-dev         # Kết nối MySQL CLI
make server-dev     # Truy cập server shell
make nginx-dev      # Truy cập nginx shell
```

### Production

```bash
make prod           # Khởi động môi trường production
make build-prod     # Build images cho production
make up-prod        # Start containers
make down-prod      # Stop containers
make restart-prod   # Restart containers
make logs-prod      # Xem logs
make db-prod        # Kết nối MySQL CLI
make server-prod    # Truy cập server shell
make nginx-prod     # Truy cập nginx shell
```

### Utility

```bash
make status         # Hiển thị trạng thái containers
make ps             # Hiển thị containers đang chạy
make clean          # Xóa tất cả containers, volumes và images
make backup-db      # Backup database
make restore-db     # Restore database
make help           # Hiển thị tất cả commands
```

## Backup & Restore Database

### Backup

```bash
make backup-db
```

File backup sẽ được lưu trong `docker/mysql/backup/` với tên format: `backup_YYYYMMDD_HHMMSS.sql`

### Restore

```bash
make restore-db
```

Chọn file backup từ danh sách hiển thị.

## Troubleshooting

### Container không khởi động

```bash
# Kiểm tra logs
make logs-dev  # hoặc logs-prod

# Kiểm tra trạng thái
make status
```

### Port đã được sử dụng

Kiểm tra và dừng service đang sử dụng port:

```bash
sudo lsof -i :80
sudo lsof -i :3306
sudo lsof -i :8080
```

### Reset môi trường

```bash
make clean
make dev  # hoặc prod
```

## Notes

- Development environment sử dụng volume mapping để hỗ trợ hot-reload
- Production environment sử dụng optimized build với binary đã compile
- Tất cả containers đều có health checks
- MySQL data được persist qua volumes
- Nginx logs được lưu trong volumes

