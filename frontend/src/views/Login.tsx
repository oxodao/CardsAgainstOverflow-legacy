import React from 'react';

import styles from '../assets/scss/login.module.scss';

export default function Login() {
    return <form className={styles.Login}>
        <label htmlFor="username">Username:</label>
        <input type="text"/>

        <input type="submit" />
    </form>;
}