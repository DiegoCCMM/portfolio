package com.adidas.backend.prioritysaleservice.models;

import java.util.Date;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonProperty;

public class AdiClubClient extends Client {
    private int points;
    private Date registrationDate;
    
    @JsonCreator
    public AdiClubClient(
            @JsonProperty("email") String emailAddress,
            @JsonProperty("points") int points,
            @JsonProperty("registrationDate") Date registrationDate) {
        super(emailAddress);
        this.points = points;
        this.registrationDate = registrationDate;
    }

    public int getPoints() {
        return points;
    }

    public void setPoints(int points) {
        this.points = points;
    }

    public Date getRegistrationDate() {
        return registrationDate;
    }

    public void setRegistrationDate(Date registrationDate) {
        this.registrationDate = registrationDate;
    }
    @Override
    public int compareTo(Client that) {
        if (that instanceof AdiClubClient){
            return this.points > ((AdiClubClient) that).points
                    || (this.points == ((AdiClubClient) that).points
                    && this.registrationDate.before(((AdiClubClient) that).registrationDate)) ? 1 : -1;
        } else {
            return 1;
        }
    }

    @Override
    public boolean isGreaterThan(Client that) {
        if (that instanceof AdiClubClient){
            return this.points > ((AdiClubClient) that).points
                    || (this.points == ((AdiClubClient) that).points
                    && this.registrationDate.before(((AdiClubClient) that).registrationDate));
        } else {
            return true;
        }
    }
    @Override
    public boolean isAdiClient(){return true;}

    @Override
    public boolean isEmailCorrect() {
        return super.getEmail().matches("[\\w-\\.]+@adiclub.com");
    }

    @Override
    public String toString() {
        return "adiClubClient{" +
                "email=" + getEmail() +
                ", points=" + points +
                ", registrationDate=" + registrationDate +
                '}';
    }
}
