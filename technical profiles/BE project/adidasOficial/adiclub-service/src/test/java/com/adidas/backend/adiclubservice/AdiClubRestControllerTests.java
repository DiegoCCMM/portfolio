package com.adidas.backend.adiclubservice;

import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.extension.ExtendWith;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.ValueSource;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.test.context.junit.jupiter.SpringExtension;

import java.time.Instant;
import java.util.Objects;

import com.adidas.backend.adiclubservice.controller.AdiClubRestController;
import com.adidas.backend.adiclubservice.dto.AdiClubMemberInfoDto;

@ExtendWith(SpringExtension.class)
@WebMvcTest(value = AdiClubRestController.class)
class AdiClubRestControllerTests {

	@Autowired
	private AdiClubRestController adiClubRestController;

	@ParameterizedTest
	@ValueSource(strings =
			{"user1@gmail.com", "user2@gmail.com", "user3@adiclub.com", "user4@adiclub.com"})
	void getAdiClubMemberInfoTest(String email) {
		ResponseEntity<AdiClubMemberInfoDto> response = adiClubRestController.getAdiClubMemberInfo(email);
		Assertions.assertEquals(HttpStatus.OK, response.getStatusCode());
		Assertions.assertInstanceOf(Integer.class, Objects.requireNonNull(response.getBody()).getPoints());
		Assertions.assertInstanceOf(Instant.class, response.getBody().getRegistrationDate());
		Assertions.assertEquals(email, response.getBody().getEmail());
	}

}
