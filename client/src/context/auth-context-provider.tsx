import { useEffect, useState, ReactNode } from 'react';
import { API_BASE_URL } from '../config';
import { AuthContext } from './auth-context';

export function AuthProvider({ children }: { children: ReactNode }) {
    const [isAuthenticated, setIsAuthenticated] = useState(false);
    const [email, setEmail] = useState<string | null>(null);
    const [userId, setUserId] = useState<string | null>(null);
    const [isLoading, setIsLoading] = useState(true);

    const checkAuth = async () => {
        try {
            const response = await fetch(`${API_BASE_URL}/auth/session`, {
                credentials: 'include',
            });
            const data = await response.json();

            setIsAuthenticated(data.authenticated);
            setEmail(data.email || null);
            setUserId(data.user_id || null);
        } catch (error) {
            console.error('Error during authentication check:', error);
            setIsAuthenticated(false);
            setEmail(null);
            setUserId(null);
        } finally {
            setIsLoading(false);
        }
    };

    const signin = async () => {
        window.location.href = `${API_BASE_URL}/auth`;
    };

    const signout = async () => {
        try {
            await fetch(`${API_BASE_URL}/auth/signout`, {
                method: 'POST',
                credentials: 'include',
            });
        } finally {
            setIsAuthenticated(false);
            setEmail(null);
            setUserId(null);
        }
    };

    useEffect(() => {
        checkAuth();
    }, []);

    return (
        <AuthContext.Provider
            value={{
                isAuthenticated,
                email,
                userId,
                isLoading,
                checkAuth,
                signin,
                signout
            }}
        >
            {children}
        </AuthContext.Provider>
    );
}

// export function useAuth() {
//     const context = useContext(AuthContext);
//     if (!context) {
//         throw new Error('useAuth must be used within an AuthProvider');
//     }
//     return context;
// }