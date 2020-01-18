CREATE DATABASE IF NOT EXISTS `kirenbbq`;

USE `kirenbbq`;

DROP TABLE IF EXISTS `products`;
SET character_set_client = utf8mb4 ;
CREATE TABLE `products` (
  `product_id` varchar(10) DEFAULT NULL,
  `product_name_th` varchar(200) DEFAULT NULL,
  `product_name_en` varchar(200) DEFAULT NULL,
  `product_description_th` varchar(1000) DEFAULT NULL,
  `product_description_en` varchar(1000) DEFAULT NULL,
  `product_image_path` varchar(200) DEFAULT NULL,
  `best_seller_flag` varchar(1) DEFAULT NULL,
  `product_price` double DEFAULT NULL,
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `products` WRITE;
INSERT INTO `products` VALUES ('AA900','เตาย่างหิน ขนาด 31×42 ซม.','','รหัสและชื่อสินค้า\\nAAA01 เตาย่างแก๊ส + แผ่นย่างชีส-ไข่ ขนาด 350 มม. (เดิมรุ่น N1)\\nในชุดประกอบด้วย:\\n1. เตาย่างแก๊สกระป๋อง ชนิดเปลว (950 บาท)\\n2. กรอบสเตนเลสสำหรับเตาแก๊ส (ไม่มีค่าใช้จ่าย)\\n3. แผ่นย่างชีส-ไข่ เคลือบ Non-stick ขนาด 350 มม. (4,500 บาท)\\n\\nสำหรับแผ่นย่าง แผ่นย่างตรงกลางสามารถเปลี่ยนเป็น:\\n– แผ่นย่างแบบมีรู\\n– แผ่นย่างแบบไม่มีรู\\n– หม้อผัด\\n– หม้อต้ม\\n*สามารถซื้อแผ่นตรงกลางเพิ่มได้ในราคาแผ่นละ 1,500 บาท\\n\\nราคารวมทั้งชุด (เตา+แผ่นย่าง): 5,450 บาท','','https://d28uedxc25vrgt.cloudfront.net/images/products/AA900.png','Y',6000,'2019-12-01 08:16:04');
INSERT INTO `products` VALUES ('AA901','เตาย่างแก๊ส กระทะหินแกรนิตขนาด 590×480 มม.','','รหัสและชื่อสินค้า\\nAAA01 เตาย่างแก๊ส + แผ่นย่างชีส-ไข่ ขนาด 350 มม. (เดิมรุ่น N1)\\nในชุดประกอบด้วย:\\n1. เตาย่างแก๊สกระป๋อง ชนิดเปลว (950 บาท)\\n2. กรอบสเตนเลสสำหรับเตาแก๊ส (ไม่มีค่าใช้จ่าย)\\n3. แผ่นย่างชีส-ไข่ เคลือบ Non-stick ขนาด 350 มม. (4,500 บาท)\\n\\nสำหรับแผ่นย่าง แผ่นย่างตรงกลางสามารถเปลี่ยนเป็น:\\n– แผ่นย่างแบบมีรู\\n– แผ่นย่างแบบไม่มีรู\\n– หม้อผัด\\n– หม้อต้ม\\n*สามารถซื้อแผ่นตรงกลางเพิ่มได้ในราคาแผ่นละ 1,500 บาท\\n\\nราคารวมทั้งชุด (เตา+แผ่นย่าง): 5,450 บาท','','https://d28uedxc25vrgt.cloudfront.net/images/products/AA901.png','Y',45000,'2019-12-01 08:16:05');
INSERT INTO `products` VALUES ('AA902','เตาย่างแก๊ส กระทะหินขนาด 700×600 มม.','','รหัสและชื่อสินค้า\\nAAA01 เตาย่างแก๊ส + แผ่นย่างชีส-ไข่ ขนาด 350 มม. (เดิมรุ่น N1)\\nในชุดประกอบด้วย:\\n1. เตาย่างแก๊สกระป๋อง ชนิดเปลว (950 บาท)\\n2. กรอบสเตนเลสสำหรับเตาแก๊ส (ไม่มีค่าใช้จ่าย)\\n3. แผ่นย่างชีส-ไข่ เคลือบ Non-stick ขนาด 350 มม. (4,500 บาท)\\n\\nสำหรับแผ่นย่าง แผ่นย่างตรงกลางสามารถเปลี่ยนเป็น:\\n– แผ่นย่างแบบมีรู\\n– แผ่นย่างแบบไม่มีรู\\n– หม้อผัด\\n– หม้อต้ม\\n*สามารถซื้อแผ่นตรงกลางเพิ่มได้ในราคาแผ่นละ 1,500 บาท\\n\\nราคารวมทั้งชุด (เตา+แผ่นย่าง): 5,450 บาท','','https://d28uedxc25vrgt.cloudfront.net/images/products/AA902.png','Y',52000,'2019-12-01 08:16:05');
INSERT INTO `products` VALUES ('AA903','เตาไฟฟ้าอินฟาเรดเพื่อการพาณิชย์ แบบเหลี่ยม-ฝัง ขนาด 2000 วัตต์','','','','https://d28uedxc25vrgt.cloudfront.net/images/products/AA903.png','Y',7000,'2019-12-01 08:16:06');
INSERT INTO `products` VALUES ('AA904','เตาไฟฟ้าอินฟาเรดเพื่อการพาณิชย์ แบบกลม-ฝัง ขนาด 2000 วัตต์','','','','https://d28uedxc25vrgt.cloudfront.net/images/products/AA904.png','Y',7000,'2019-12-01 08:16:07');
INSERT INTO `products` VALUES ('AAA01','เตาย่างแก๊ส + แผ่นย่างชีส-ไข่ ขนาด 350 มม.รุ่น N 1','','','','https://d28uedxc25vrgt.cloudfront.net/images/products/AAA01.png','Y',5450,'2019-12-01 08:16:07');
INSERT INTO `products` VALUES ('AAA02','ชุดเตาย่างแก๊ส + แผ่นย่างเนื้อ-ซุปเคลือบ Non-stick ขนาด 310 มม.','','','','https://d28uedxc25vrgt.cloudfront.net/images/products/AAA02.png','Y',2400,'2019-12-01 08:16:08');
INSERT INTO `products` VALUES ('AAA03','ชุดเตาย่างแก๊ส + แผ่นย่าง + หม้อต้ม + ฝาแก้ว','','','','https://d28uedxc25vrgt.cloudfront.net/images/products/AAA03.png','Y',3300,'2019-12-01 08:16:09');
INSERT INTO `products` VALUES ('AAA06','ชุดเตาย่างแก๊สกระป๋อง + กระทะย่างรุ่น 14','','','','https://d28uedxc25vrgt.cloudfront.net/images/products/AAA06.png','N',2950,'2019-12-01 08:16:09');
INSERT INTO `products` VALUES ('AAA07','ชุดเตาย่างแก๊สกระป๋อง + กระทะย่างรุ่น 13','','','','https://d28uedxc25vrgt.cloudfront.net/images/products/AAA07.png','N',2950,'2019-12-01 08:16:10');
INSERT INTO `products` VALUES ('AAA09','ชุดเตาย่างแก๊สกระป๋อง + กระทะย่างชีสชาบู','','','','https://d28uedxc25vrgt.cloudfront.net/images/products/AAA09.png','N',5450,'2019-12-01 08:16:11');
UNLOCK TABLES;
