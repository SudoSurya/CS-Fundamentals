- @Component
![image](https://github.com/user-attachments/assets/49fd838a-df92-42fb-bc7a-a65a46114e34)

- @Autowired
![image](https://github.com/user-attachments/assets/278683a0-30fb-4c8e-ae26-64893dde4fe7)

- @Primary
![image](https://github.com/user-attachments/assets/e75cdfd0-c385-45cb-9811-64617b693540)

- @Qualifier

![image](https://github.com/user-attachments/assets/106ed138-7c1f-4a8c-9f82-af18f11cdca4)

- Field Injection 
![image](https://github.com/user-attachments/assets/f0a65642-ae44-4702-a465-b820f6acf9c1)
- Setter Injection 
![image](https://github.com/user-attachments/assets/130facc0-cb59-4789-89aa-b77cde12aa4e)
- Constructor Injection 
![image](https://github.com/user-attachments/assets/d41c28ed-8f71-4809-961a-ef66f0d8c88f)

## Recommendations

| Type        | When to Use |
|------------|----------------------------------------------------------|
| **Constructor** | Default choice for mandatory, immutable dependencies. |
| **Setter**      | Rarely, for optional/mutable dependencies.            |
| **Field**       | Avoid (except in legacy code or frameworks requiring it). |

## Key Takeaways

- **Constructor injection** is strongly recommended by modern frameworks (Spring, Jakarta EE) for its safety and clarity.
- **Avoid field injection** in most cases—it’s less visible and harder to test.
- **Use setter injection** only when dependencies need to change post-creation (rare).
