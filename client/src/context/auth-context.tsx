import { createContext } from 'react';

interface AuthContextType {
    isAuthenticated: boolean;
    email: string | null;
    userId: string | null;
    isLoading: boolean;
    checkAuth: () => Promise<void>;
    signin: () => Promise<void>;
    signout: () => Promise<void>;
}

export const AuthContext = createContext<AuthContextType | null>(null);
