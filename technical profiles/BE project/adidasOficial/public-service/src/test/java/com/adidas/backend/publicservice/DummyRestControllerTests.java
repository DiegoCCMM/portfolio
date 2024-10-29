package com.adidas.backend.publicservice;

import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.test.context.junit.jupiter.SpringExtension;

import com.adidas.backend.publicservice.controller.PublicRestController;

@ExtendWith(SpringExtension.class)
@WebMvcTest(value = PublicRestController.class)
class DummyRestControllerTests {

	@Autowired
	private PublicRestController dummyRestController;

	@Test
	void testDummyController() {
		String expectedResponse = "Hello, this is a dummy response from public service";
		ResponseEntity<String> response = dummyRestController.getDummyEndpointResponse();
		Assertions.assertEquals(HttpStatus.OK, response.getStatusCode());
		Assertions.assertEquals(expectedResponse, response.getBody());
	}

}
