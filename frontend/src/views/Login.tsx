import React from 'react';

import styles from '../assets/scss/login.module.scss';

export default function Login() {
    //const [formData, setFormData] = useForm();

    return <form className={styles.Login}>
        <label htmlFor="username">Username:</label>
        <input type="text"/>

        <input type="submit" />
    </form>;
}