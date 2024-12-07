import { useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { useAuth } from '@/context/auth-context';

export function ProtectedRoute({ children }: { children: React.ReactNode }) {
    const { isAuthenticated, isLoading } = useAuth();
    const navigate = useNavigate();

    useEffect(() => {
        if (!isLoading && !isAuthenticated) {
            navigate('/');
        }
    }, [isLoading, isAuthenticated, navigate]);

    if (isLoading) {
        return <div>Loading...</div>;
    }

    return isAuthenticated ? <>{children}</> : null;
}