# E-Commerce Project Documentation

## Modules - 70% OFF

### MODULE 1 - Admin Home

- **Login / Logout:**
  - Admin can log in using predefined username and password.
  - Admin can end the session using the logout button.

- **Dashboard:**
  - Overview of sales and revenue using graphs.

### MODULE 2 - Users Home

- **User Login:**
  - Users need to log in with their credentials to buy from the store.

- **User Signup:**
  - Users can become a member of the store using the signup option.
  - Provide credentials for later login.

- **Browse Product:**
  - Anyone can browse and see the products available in the store.

- **Browse Category:**
  - Products are split into categories and subcategories/brands.

### MODULE 3 - User Product

- **Product Details:**
  - Displays detailed information, images, and zoom views of the product.
  - Description of the product.

- **Add to Cart:**
  - Users can add products to the cart before checkout.

- **Add to Wishlist:**
  - Users can favorite a product and add it to their wishlist.

### MODULE 4 - Admin Product Management

- **Add Product:**
  - Admin can add new products to the database.

- **Edit Product / Add Stock:**
  - Admin can update the products.
  - Manage stock information.

- **Search for Product:**
  - Admin can search for products to view/update them.

- **Remove Product:**
  - Admin can delete a product from the database.

### MODULE 5 - Admin User Management

- **Delete User:**
  - Admin can delete a user.

- **Block User:**
  - Admin can ban a user from the store.

- **Search User:**
  - Admin can search for users and view their information.

### MODULE 6 - Admin Order and Payment

- **Order Status:**
  - Admin can change the user order status (Confirmed, Delivered).

- **Cancel Order:**
  - Admin can cancel user orders if they are pending.

- **Payment Methods:**
  - Admin can manage the payment methods and details of the store.

### MODULE 7 - Profile

- **Change Password:**
  - Users can change their password and set a new one.

- **Edit Details:**
  - Users can change their address and personal information.

- **Order History:**
  - Users can view their purchase/order history.

### MODULE 8 - User Order and Payment

- **Checkout:**
  - Users can go to checkout directly from their cart and see order details.

- **Address:**
  - Users can select their address or add a new address for delivery.

- **Payment Methods:**
  - Users can select their payment method and proceed to the Payment Gateway.

### MODULE 9 - Sales Report

- **Sales Report:**
  - Report on all sales with the option to set the range.

- **Print Sales Report:**
  - Convert sales report into PDF and Excel sheet.

---

# API Specification Doc

## Version: 1.0

- **Date:** 22-04-2021
- **Author:** ROHITH ER
- **Description:** Initial draft

### 1. ADMIN:

#### 1.1 Admin Login:

- **Method:** POST
- **URL:** /admin_login

**Request:**

- **Parameter Name:**
  - User Name (Type & Length: String)
  - Password (Type & Length: Password)

**Response:**

- **Parameter Name:**
  - 200 (Type & Length: "Successfully Login to Admin panel")
  - 501 (Type & Length: {"error": "Username and password invalid"})

#### 1.2 Admin Panel:

- **Method:** GET
- **URL:** /admin_panel

**Response:**

- **Parameter Name:**
  - 200 (Type & Length: 'Success'. Enter to admin_Panel)

#### 1.3 PRODUCTS:

#### 1.3.1 View Products:

- **Method:** GET
- **URL:** /admin_panel/products

**Response:**

- **Parameter Name:**
  - 200 (Type & Length: OK)

#### 1.3.2 Add Products:

- **Method:** POST
- **URL:** /admin_panel/products/add_product

**Request:**

- **Parameter Name:**
  - Category (Type & Length: String)
  - Product name (Type & Length: String)
  - Price (Type & Length: int)
  - Quantity (Type & Length: int)
  - Product Image (Type & Length: Image)
  - Size (Type & Length: String)
  - Brand (Type & Length: String)

**Response:**

- **Parameter Name:**
  - 200 (Type & Length: 'Success')

#### 1.3.3 Edit Products:

- **Method:** PUT
- **URL:** /admin_panel/products/edit_products/id

**Request:**

- **Parameter Name:**
  - Category (Type & Length: String)
  - Product name (Type & Length: String)
  - Price (Type & Length: int)
  - Quantity (Type & Length: int)
  - Product Image (Type & Length: Image)
  - Size (Type & Length: String)
  - Brand (Type & Length: String)

**Response:**

- **Parameter Name:**
  - 200 (Type & Length: 'Success')

#### 1.3.4 Delete Products:

- **Method:** DELETE
- **URL:** /admin_panel/products/delete_products/id

**Response:**

- **Parameter Name:**
  - 200 (Type & Length: OK)

### 2. USERS:

#### 2.1 View Users:

- **Method:** GET
- **URL:** /admin_panel/user_management

**Response:**

- **Parameter Name:**
  - 200 (Type & Length: OK)

#### 2.2 Edit Users:

- **Method:** PUT
- **URL:** /admin_panel/user_management/edit_user/user_id

**Request:**

- **Parameter Name:**
  - Name (Type & Length: String)
  - Username (Type & Length: String)
  - Address (Type & Length: String)
  - Email (Type & Length: Email)

**Response:**

- **Parameter Name:**
  - 200 (Type & Length: 'Success')

#### 2.3 Delete Users:

- **Method:** DELETE
- **URL:** /admin_panel/user_management/delete_user/user_id

**Response:**

- **Parameter Name:**
  - 200 (Type & Length: OK)

### 3. ORDERS:

#### 3.1 View Orders:

- **Method:** GET
- **URL:** /admin_panel/orders

**Response:**

- **Parameter Name:**
  - 200 (Type & Length: OK)

#### 3.2 Edit Orders:

- **Method:** PUT
- **URL:** /admin_panel/orders/edit_orders/order_id

**Request:**

- **Parameter Name:**
  - username (Type & Length: String)
  - Address (Type & Length: String)
  - Product name (Type & Length: String)
  - category (Type & Length: String)
  - Quantity (Type & Length: int)
  - price (Type & Length: int)
  - size (Type & Length: String)

**Response:**

- **Parameter Name:**
  - 200 (Type & Length: 'Success')

#### 3.3 Delete Orders:

- **Method:** DELETE
- **URL:** /admin_panel/products/delete_orders/orders_id

**Response:**

- **Parameter Name:**
  - 200 (Type & Length: OK)

### 4. ADMIN LOGOUT:

- **Method:** GET
- **URL:** /admin_logout

**Response:**

- **Parameter Name:**
  - 200 (Type & Length: OK)

### 5. LANDING PAGE:

#### 5.1 Loading Lead Page:

- **Method:** GET
- **URL:** http://127.0.0.1:8000/

**Response:**

- **Parameter Name:**
  - 200 (Type & Length: OK)

### 6. USER:

#### 6.1 User Sign Up:

- **Method:** POST
- **URL:** /user_registration

**Request:**

- **Parameter Name:**
  - Name (Type & Length: String)
  - Username (Type & Length: String)
  - Address (Type & Length: String)
  - Email (Type & Length: Email)
  - Phone No. (Type & Length: int)
  - Password (Type & Length: password)

**Response:**

- **Parameter Name:**
  - 200 (Type & Length: 'Success')

#### 6.2 User Login:

- **Method:** POST
- **URL:** /user_signin

**Request:**

- **Parameter Name:**
  - Username (Type & Length: String)
  - Password (Type & Length: password)

**Response:**

- **Parameter Name:**
  - 200 (Type & Length: 'Success')
