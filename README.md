# API Documentation

## docs/openapi.yaml
**Description:** Location of yaml file above.

## Overview

This document provides an overview of the API endpoints available in the system, including authentication, user management, payments, and more.

## Base URL

https://www.example.com/api/v1

## Authentication

**Description:** Endpoints for user registration and authentication.

**Endpoint:** `POST /auth/register` **Description:** Register a new user.

**Endpoint:** `POST /auth/login` **Description:** Login a new user.

**Endpoint:** `POST /auth/social` **Description:** Authenticate a user using a social provider.

**Endpoint:** `POST /auth/magic-link` **Description:** Request a magic link for authentication.

**Endpoint:** `GET /auth/magic-link/verify` **Description:** Verify the magic link token.

**Endpoint:** `POST /auth/magic-link/password-reset` **Description:** Request a password reset link.

**Endpoint:** `POST /auth/magic-link/password-reset/verify` **Description:** Verify the password reset token and set a new password.

**Endpoint:** `PUT /auth/change-password` **Description:** Verify the password reset token and set a new password.

## User

**Description:** User related operations.

**Endpoint:** `GET /users` **Description:** Get all users.

**Endpoint:** `GET /users/{userId}` **Description:** Get single user.

**Endpoint:** `PUT /users/{userid}` **Description:** Update single user.

**Endpoint:** `DELETE /users/{userId}` **Description:** Delete single user.

**Endpoint:** `GET /users/{userId}/organisation` **Description:** Query organisations where a user belongs.

**Endpoint:** `GET /users/{userId}/data` **Description:** Retrieve user-specific data for the dashboard.

**Endpoint:** `GET /users/settings` **Description:** Retrieve user settings.

**Endpoint:** `UPDATE /users/settings` **Description:** Update user settings.

## Organisation

**Description:** Organisation related operations.

**Endpoint:** `GET /organisations` **Description:** Get all users.

**Endpoint:** `POST /organisation` **Description:** Register a new organisation.

**Endpoint:** `GET /organisations` **Description:** Get all organisation.

**Endpoint:** `GET /organisations/{orgId}` **Description:** Get an organisation by Id.

**Endpoint:** `DELETE /organisations/{orgId}` **Description:** Delete an organisation by id.

**Endpoint:** `POST /organisations/{orgId}/users` **Description:** Add a user to an organisation by id.

**Endpoint:** `PUT /organisations/{orgId}/users` **Description:** Delete a user from an organisation.

**Endpoint:** `GET /organisations/{orgId}/users` **Description:** Get users in an organisation.

## Super Admin

**Description:** Superadmin operations.

**Endpoint:** `GET /users/{userId}` **Description:** Get single user.

**Endpoint:** `PUT /users/{userId}` **Description:** Update user info.

**Endpoint:** `DELETE /users/{userId}` **Description:** Delete a user.

**Endpoint:** `GET /users/{userId}/organisation` **Description:** Query organisations where a user belongs.

**Endpoint:** `GET /users` **Description:** Get all users.

**Endpoint:** `GET /payment` **Description:** List all payments.

**Endpoint:** `GET /payments/{paymentId}` **Description:** Get payment by ID.

**Endpoint:** `PUT /payments/{paymentId}` **Description:** Update an existing payment.

## Invite

**Description:** API for user invite management.

**Endpoint:** `POST /invite/generate` **Description:** Generates an invitation token.

**Endpoint:** `POST /invite/verify` **Description:** Verify.

## Notifications

**Description:** Notification operations.

**Endpoint:** `POST /notifications/{user_id}` **Description:** Get notifications by user Id.

## Pages

**Description:** Pages available.

**Endpoint:** `POST /contact` **Description:** Submit a contact form.

**Endpoint:** `GET /contacts/messages` **Description:** Retrieve all contact messages (Admin only).

**Endpoint:** `POST /waitlist` **Description:** Add a user to the waitlist.

**Endpoint:** `GET /waitlist` **Description:** Retrieve the waitlist (Admin only).

**Endpoint:** `POST /pages/marketing` **Description:** Update marketing page content (Admin only).

**Endpoint:** `POST /data/charts` **Description:** Retrieve data for charts.

**Endpoint:** `GET /blogs` **Description:** Retrieve a list of blog posts.

**Endpoint:** `GET /blogs` **Description:** Retrieve a list of blog posts.

## Payment

**Description:** Payment operations.

**Endpoint:** `GET /payments` **Description:** List all payments.

**Endpoint:** `POST /payments` **Description:** Create a new payment.

**Endpoint:** `GET /payments/{paymentID}` **Description:** Get payment by ID.

**Endpoint:** `PUT /payments/{paymentID}` **Description:** Update an existing payment.

**Endpoint:** `POST /payments/stripe` **Description:** Create a new payment via Stripe.

**Endpoint:** `POST /payments/flutterwave` **Description:** Create a new payment via Flutterwave.

**Endpoint:** `POST /payments/lemonsqueezy` **Description:** Create a new payment via Lemonsqueezy.

<p align="left"><a href="https://hng-task-three-pt7rk.ondigitalocean.app/" target="_blank">Link to HTML version API documentation</a><p>

<p align="left"><a href="https://app.swaggerhub.com/apis/HARYORTORLARH/hng-task3/1.0.0" target="_blank">Link to swagger version API documentation</a><p>

<p align="left"><a href="https://drive.google.com/file/d/1-8mI3JqCDoQGRMwzU1yFNEq5YAdPVD6r/view" target="_blank">Link to Database Image</a><p>
