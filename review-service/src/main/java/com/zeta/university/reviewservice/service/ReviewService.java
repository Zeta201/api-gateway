package com.zeta.university.reviewservice.service;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.zeta.university.reviewservice.entity.Review;
import com.zeta.university.reviewservice.repository.ReviewRepository;

@Service
public class ReviewService {
	@Autowired
	private ReviewRepository repository;
	
	public Review submitReview(Review review) {
		return repository.save(review);
	}
	
	public List<Review> getAllReviews(){
		return repository.findAll();
	}
}
