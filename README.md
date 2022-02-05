# Account Service Layer for Melanoma Project

This is a code I contributed in the project Melanoma.

The project was founded to be an assistant platform for medical staff to diagnose a skin disease.
I participated in this project as a back-end developer and was assigned to implement an Account Service layer using GO and Fiber framework for first time.

Account layer is to connect the database for creating, updating, and deleting user accounts, authenticate a user's login requests and also send email to the users if needed.

The project is still implementing and will be public soon.

### Endpoint

- `POST` `/signin`
- `POST` `/auth`
- `GET` `/users`
- `GET` `/users/me`
- `GET` `/users/:username`
- `GET` `/users/:username/reset`
- `POST` `/password/reset?resetToken={resetToken}`

You can see what I do during the project on Medium.com via these links below.

[มือใหม่หัดทำ Backend Web Application ได้เรียนรู้อะไรบ้าง? EP.1](https://medium.com/@thanin.sawetkititham/%E0%B8%A1%E0%B8%B7%E0%B8%AD%E0%B9%83%E0%B8%AB%E0%B8%A1%E0%B9%88%E0%B8%AB%E0%B8%B1%E0%B8%94%E0%B8%97%E0%B8%B3-backend-web-application-%E0%B9%84%E0%B8%94%E0%B9%89%E0%B9%80%E0%B8%A3%E0%B8%B5%E0%B8%A2%E0%B8%99%E0%B8%A3%E0%B8%B9%E0%B9%89%E0%B8%AD%E0%B8%B0%E0%B9%84%E0%B8%A3%E0%B8%9A%E0%B9%89%E0%B8%B2%E0%B8%87-a31a305d5952)
