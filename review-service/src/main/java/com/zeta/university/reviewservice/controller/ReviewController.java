package com.zeta.university.reviewservice.controller;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.zeta.university.reviewservice.common.ReviewRequest;
import com.zeta.university.reviewservice.entity.Review;
import com.zeta.university.reviewservice.service.ReviewService;

@RestController
@RequestMapping("/api/v1/review-service")
public class ReviewController {
	
	@Autowired
	private ReviewService service;
	
	@PostMapping("/submit")
	public Review submitReview(@RequestBody ReviewRequest request) {
		Review newReview = Review.builder().username(request.getUsername())
				.content(request.getContent()).stars(request.getStars()).build();
		return service.submitReview(newReview);
	}
	
	@GetMapping("/")
	public List<Review> getAllReviews() {
		
		return service.getAllReviews();
	}
}
