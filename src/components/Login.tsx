import React, { useState } from 'react';

interface LoginButtonProps {
  loading?: boolean;
}

export function LoginButton({ loading = false }: LoginButtonProps) {
  const [hovered, setHovered] = useState(false);
  const disabled = loading;
  return (
    <button
      type="button"
      title={loading ? 'Signing in…' : 'Sign in to your account'}
      aria-label={loading ? 'Signing in' : 'Sign in to your account'}
      aria-busy={loading}
      disabled={disabled}
      data-testid="login-button"
      onMouseEnter={() => setHovered(true)}
      onMouseLeave={() => setHovered(false)}
      style={{
        display: 'inline-flex',
        alignItems: 'center',
        gap: 8,
        padding: '10px 18px',
        borderRadius: 8,
        border: 'none',
        cursor: disabled ? 'not-allowed' : 'pointer',
        opacity: disabled ? 0.7 : 1,
        fontSize: 16,
        fontWeight: 600,
        color: 'white',
        background: hovered && !disabled ? '#1d4ed8' : '#2563eb',
        transition: 'background 0.15s ease, opacity 0.15s ease',
      }}
    >
      {loading ? (
        <svg
          data-testid="login-button-spinner"
          width="18"
          height="18"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          strokeWidth="2"
          strokeLinecap="round"
          aria-hidden="true"
          style={{
            animation: 'proctor-spin 0.8s linear infinite',
          }}
        >
          <path d="M12 2 A 10 10 0 0 1 22 12" />
        </svg>
      ) : (
        <svg
          data-testid="login-button-icon"
          width="18"
          height="18"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          strokeWidth="2"
          strokeLinecap="round"
          strokeLinejoin="round"
          aria-hidden="true"
        >
          <rect x="5" y="11" width="14" height="9" rx="2"></rect>
          <path d="M8 11V8a4 4 0 0 1 8 0v3"></path>
        </svg>
      )}
      <style>{`
        @keyframes proctor-spin {
          from { transform: rotate(0deg); }
          to   { transform: rotate(360deg); }
        }
      `}</style>
      {loading ? 'Signing in…' : 'Sign in'}
    </button>
  );
}
