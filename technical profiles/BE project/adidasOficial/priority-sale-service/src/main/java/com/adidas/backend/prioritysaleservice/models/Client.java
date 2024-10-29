package com.adidas.backend.prioritysaleservice.models;

public class Client implements Comparable<Client> {
    private String email;

    public Client(String emailAddress) {
        this.email = emailAddress;
    }

    public String getEmail() {
        return email;
    }

    public void setEmail(String emailAddress) {
        this.email = emailAddress;
    }

    public int compareTo(Client that) {
        return -1;
    }

    public boolean isGreaterThan(Client that) {
        return !(that instanceof AdiClubClient);
    }

    public boolean isAdiClient(){return false;}

    public boolean isEmailCorrect() {
        return email.matches("[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}");
    }

    @Override
    public String toString() {
        return "client{" +
                "email='" + email + '\'' +
                '}';
    }
}
